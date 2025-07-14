package analyzer

import (
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

    return langCount
}
