package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/Yasaswini-Devi/git-stalker/api"
    "github.com/Yasaswini-Devi/git-stalker/analyzer"
    "github.com/Yasaswini-Devi/git-stalker/report"
)

var rootCmd = &cobra.Command{
    Use:   "git-stalker [username]",
    Short: "GitHub developer profiler",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        username := args[0]
        fmt.Println("🔍 Profiling:", username)

        name, bio := api.FetchUserProfile(username)
        repos := api.FetchUserRepos(username)
        if len(repos) == 0 {
            fmt.Println("❌ No repos found or error occurred.")
            return
        }

        languageMap := analyzer.AnalyzeLanguages(repos)
        hourCount, dayCount, totalCommits, archetype := analyzer.AnalyzeCommitActivity(username, repos)

        err := report.GenerateMarkdownReport(username, name, bio, languageMap, hourCount, dayCount, archetype, totalCommits)
        if err != nil {
            fmt.Println("⚠️ Failed to generate markdown report:", err)
        }
    },
}

func Execute() {
    cobra.CheckErr(rootCmd.Execute())
}

