package cmd

import (
	"fmt"
	"io"
	"log"
	"time"

	"github.com/Yasaswini-Devi/git-stalker/analyzer"
	"github.com/Yasaswini-Devi/git-stalker/api"
	"github.com/Yasaswini-Devi/git-stalker/report"
	"github.com/spf13/cobra"
)

var generateMarkdown bool
var openReport bool
var verboseLog bool

var rootCmd = &cobra.Command{
	Use:   "git-stalker [username]",
	Short: "GitHub developer profiler",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if !verboseLog {
			log.SetOutput(io.Discard)
		}

		username := args[0]
		log.Println("🔍 Profiling GitHub user:", username)

		name, bio := api.FetchUserProfile(username)
		repos := api.FetchUserRepos(username)

		fmt.Println("\n👤 Name:", name)
		fmt.Println("📝 Bio:", bio)
		if len(repos) == 0 {
			log.Println("❌ No repositories found or an error occurred.")
			return
		}
		fmt.Printf("Public Repos: %v\n", len(repos))

		languageStats := analyzer.AnalyzeLanguages(repos)
		fmt.Println("\n💻 Language Usage:")
		for lang, count := range languageStats {
			fmt.Printf("• %-15s : %d repos\n", lang, count)
		}

		commitsByHour, commitsByDay, totalCommits, archetype := analyzer.AnalyzeCommitActivity(username, repos)
		fmt.Println("\nTotal Commits: ", totalCommits)
		fmt.Printf("\n⏰ Commits by Hour:\n")
		for hour := 0; hour < 24; hour++ {
			fmt.Printf("• %02d:00 – %d commits\n", hour, commitsByHour[hour])
		}

		fmt.Printf("\n📅 Commits by Day:\n")
		days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
		for i, day := range days {
			weekday := time.Weekday(i)
			fmt.Printf("• %-9s – %d commits\n", day, commitsByDay[weekday])
		}

		fmt.Printf("\n🧠 Developer Archetype: %s\n", archetype)

		topRepos := analyzer.TopStarredRepos(repos, 3)
		fmt.Println("\n🏆 Top Starred Repositories:")
		for _, repo := range topRepos {
			fmt.Printf("• %s – ⭐ %d stars\n", repo.Name, repo.StargazersCount)
		}

		if generateMarkdown {
			err := report.GenerateMarkdownReport(username, name, bio, languageStats, commitsByHour, commitsByDay, archetype, totalCommits, topRepos)
			if err != nil {
				log.Printf("⚠️ Failed to generate markdown report: %v\n", err)
			} else if openReport {
				report.OpenMarkdownReport(username)
			}
		}
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&generateMarkdown, "md", "m", false, "Generate markdown report")
	rootCmd.Flags().BoolVarP(&openReport, "open", "o", false, "Open markdown report after generation")
	rootCmd.Flags().BoolVarP(&verboseLog, "log", "l", false, "Enable detailed logging output")
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
