package report

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Yasaswini-Devi/git-stalker/api"
)

// GenerateMarkdownReport creates a developer profile report in markdown format.
func GenerateMarkdownReport(
	username, name, bio string,
	langs map[string]int,
	hourMap map[int]int,
	dayMap map[time.Weekday]int,
	archetype string,
	totalCommits int,
	topRepos []api.Repo,
) error {
	var sb strings.Builder

	badge := GenerateBadgeSnippet(archetype)

	// Header
	sb.WriteString(fmt.Sprintf("# GitHub Profile: [%s](https://github.com/%s)\n\n", username, username))
	sb.WriteString(fmt.Sprintf("%s\n\n", badge))
	sb.WriteString(fmt.Sprintf("**Name**: %s\n\n", name))
	sb.WriteString(fmt.Sprintf("**Bio**: %s\n\n", bio))
	sb.WriteString(fmt.Sprintf("**Developer Archetype**: %s\n\n", archetype))
	sb.WriteString(fmt.Sprintf("**Total Commits Analyzed**: %d\n\n", totalCommits))

	// Language Usage
	sb.WriteString("## 💻 Language Usage\n")
	for lang, count := range langs {
		sb.WriteString(fmt.Sprintf("- **%s**: %d repos\n", lang, count))
	}

	// Commits by Hour
	sb.WriteString("\n## ⏰ Commits by Hour\n")
	for hour := 0; hour < 24; hour++ {
		sb.WriteString(fmt.Sprintf("- %02d:00 → %d commits\n", hour, hourMap[hour]))
	}

	// Commits by Day
	sb.WriteString("\n## 📅 Commits by Day\n")
	for d := time.Sunday; d <= time.Saturday; d++ {
		sb.WriteString(fmt.Sprintf("- %s → %d commits\n", d.String(), dayMap[d]))
	}

	// Top Starred Repositories
	sb.WriteString("\n## 🏆 Top Starred Repositories\n")
	if len(topRepos) == 0 {
		sb.WriteString("- No public repos with stars found.\n")
	} else {
		for _, repo := range topRepos {
			sb.WriteString(fmt.Sprintf("- [%s](https://github.com/%s/%s) — ⭐ %d stars\n",
				repo.Name, username, repo.Name, repo.StargazersCount))
		}
	}

	// Save to file
	filename := fmt.Sprintf("%s_report.md", username)
	if err := os.WriteFile(filename, []byte(sb.String()), 0644); err != nil {
		return fmt.Errorf("failed to write markdown report: %w", err)
	}

	fmt.Println("📄 Markdown report saved as:", filename)
	return nil
}
