package analyzer

import (
    "fmt"
    "github.com/Yasaswini-Devi/git-stalker/api"
)

func AnalyzeLanguages(repos []api.Repo) {
    langCount := make(map[string]int)

    for _, repo := range repos {
        if repo.Language != "" {
            langCount[repo.Language]++
        }
    }

    fmt.Println("\n💻 Language Usage:")
    for lang, count := range langCount {
        fmt.Printf("• %-15s : %d repos\n", lang, count)
    }
}
