package analyzer

import (
    "fmt"
    "github.com/Yasaswini-Devi/git-stalker/api"
)

func AnalyzeLanguages(repos []api.Repo) map[string]int {
    langCount := make(map[string]int)
    for _, repo := range repos {
        lang := repo.Language
        if lang != "" {
            langCount[lang]++
        }
    }

    fmt.Println("\n💻 Language Usage:")
    for lang, count := range langCount {
        fmt.Printf("• %-15s : %d repos\n", lang, count)
    }

    return langCount
}
