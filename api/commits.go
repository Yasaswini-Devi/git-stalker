package api

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type Commit struct {
	CommitData struct {
		Author struct {
			Date time.Time `json:"date"`
		} `json:"author"`
	} `json:"commit"`
}

// FetchCommitTimestamps returns commit timestamps for a given repo of a user.
func FetchCommitTimestamps(username, repoName string) []time.Time {
	url := "https://api.github.com/repos/" + username + "/" + repoName + "/commits?per_page=30"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("❌ Failed to create request: %v", err)
		return nil
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("GITHUB_TOKEN"))
	req.Header.Set("Accept", "application/vnd.github+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("❌ Failed to fetch commits for %s/%s: %v", username, repoName, err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("⚠️  Skipping %s/%s (status: %d)", username, repoName, resp.StatusCode)
		return nil
	}

	var commits []Commit
	if err := json.NewDecoder(resp.Body).Decode(&commits); err != nil {
		log.Printf("❌ Failed to parse commits for %s/%s: %v", username, repoName, err)
		return nil
	}

	var commitTimes []time.Time
	for _, commit := range commits {
		commitTimes = append(commitTimes, commit.CommitData.Author.Date)
	}

	return commitTimes
}
