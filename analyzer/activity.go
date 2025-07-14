package analyzer

import (
	"github.com/Yasaswini-Devi/git-stalker/api"
	"time"
)

func AnalyzeCommitActivity(username string, repos []api.Repo) (map[int]int, map[time.Weekday]int, int, string) {
	hourCount := make(map[int]int)
	dayCount := make(map[time.Weekday]int)
	totalCommits := 0

	for _, repo := range repos {
		timestamps := api.FetchCommitTimestamps(username, repo.Name)
		for _, ts := range timestamps {
			hourCount[ts.Hour()]++
			dayCount[ts.Weekday()]++
			totalCommits++
		}
	}
	archetype := GetArchetype(hourCount, dayCount)

	return hourCount, dayCount, totalCommits, archetype
}
