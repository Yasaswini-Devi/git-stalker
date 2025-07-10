package analyzer

import (
    "sort"
    "github.com/Yasaswini-Devi/git-stalker/api"
)

func TopStarredRepos(repos []api.Repo, limit int) []api.Repo {
    sorted := make([]api.Repo, len(repos))
    copy(sorted, repos)

    sort.Slice(sorted, func(i, j int) bool {
        return sorted[i].StargazersCount > sorted[j].StargazersCount
    })

    if len(sorted) > limit {
        return sorted[:limit]
    }
    return sorted
}
