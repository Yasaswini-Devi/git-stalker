package report

import (
    "fmt"
    "strings"
)

func GenerateBadgeSnippet(archetype string) string {
    var safeLabel string

    switch {
    case strings.Contains(archetype, "Night Owl"):
        safeLabel = "Night%20Owl"
    case strings.Contains(archetype, "Weekend Hacker"):
        safeLabel = "Weekend%20Hacker"
    case strings.Contains(archetype, "9-to-5"):
        safeLabel = "9to5%20Coder" // Remove dashes to avoid issues
    case strings.Contains(archetype, "Chaos"):
        safeLabel = "Chaos%20Coder"
    default:
        safeLabel = "Developer"
    }

    badgeURL := fmt.Sprintf("https://img.shields.io/badge/Archetype-%s-blueviolet", safeLabel)
    markdown := fmt.Sprintf("![Dev Archetype: %s](%s)", archetype, badgeURL)
    return markdown
}
