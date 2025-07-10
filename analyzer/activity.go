package analyzer

import (
    "fmt"
    "github.com/Yasaswini-Devi/git-stalker/api"
    "time"
)

func AnalyzeCommitActivity(username string, repos []api.Repo) {
    hourCount := make(map[int]int)
    dayCount := make(map[time.Weekday]int)

    for _, repo := range repos {
        times := api.FetchCommitTimestamps(username, repo.Name)
        for _, t := range times {
            hourCount[t.Hour()]++
            dayCount[t.Weekday()]++
        }
    }

    fmt.Println("\n⏰ Coding Activity by Hour:")
    for h := 0; h < 24; h++ {
        fmt.Printf("%02d:00 - %2d commits\n", h, hourCount[h])
    }

    fmt.Println("\n📅 Coding Activity by Day:")
    for d := time.Sunday; d <= time.Saturday; d++ {
        fmt.Printf("%-9s: %2d commits\n", d.String(), dayCount[d])
    }

        archetype := GetArchetype(hourCount, dayCount)
    fmt.Println("\n🧠 Developer Archetype:", archetype)
}
