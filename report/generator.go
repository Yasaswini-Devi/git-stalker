package report

import (
    "fmt"
    "os"
    "strings"
    "time"
    "github.com/Yasaswini-Devi/git-stalker/api"
)

func GenerateMarkdownReport(username, name, bio string, langs map[string]int, hourMap map[int]int, dayMap map[time.Weekday]int, archetype string, totalCommits int, topRepos []api.Repo) error {
    var sb strings.Builder

    badge := GenerateBadgeSnippet(archetype)
    
    sb.WriteString(fmt.Sprintf("# GitHub Profile: [%s](https://github.com/%s)\n\n", username, username))
    sb.WriteString(fmt.Sprintf("%s\n\n", badge))
    sb.WriteString(fmt.Sprintf("**Name**: %s\n\n", name))
    sb.WriteString(fmt.Sprintf("**Bio**: %s\n\n", bio))
    sb.WriteString(fmt.Sprintf("**Developer Archetype**: %s\n\n", archetype))
    sb.WriteString(fmt.Sprintf("**Total Commits Analyzed**: %d\n\n", totalCommits))

    sb.WriteString("## 💻 Language Usage\n")
    for lang, count := range langs {
        sb.WriteString(fmt.Sprintf("- **%s**: %d repos\n", lang, count))
    }

    sb.WriteString("\n## ⏰ Commits by Hour\n")
    for hour := 0; hour < 24; hour++ {
        sb.WriteString(fmt.Sprintf("- %02d:00 → %d commits\n", hour, hourMap[hour]))
    }

    sb.WriteString("\n## 📅 Commits by Day\n")
    for d := time.Sunday; d <= time.Saturday; d++ {
        sb.WriteString(fmt.Sprintf("- %s → %d commits\n", d.String(), dayMap[d]))
    }

    sb.WriteString("\n## 🏆 Top Starred Repositories\n")
    for _, repo := range topRepos {
        sb.WriteString(fmt.Sprintf("- %s — ⭐ %d stars\n", repo.Name, repo.StargazersCount))
    }

    filename := fmt.Sprintf("%s_report.md", username)
    err := os.WriteFile(filename, []byte(sb.String()), 0644)
    if err != nil {
        return err
    }

    fmt.Println("📄 Markdown report saved as:", filename)
    return nil
}
