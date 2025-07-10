package api

import (
    "encoding/json"
    "fmt"
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

func FetchCommitTimestamps(username, repoName string) []time.Time {
    url := fmt.Sprintf("https://api.github.com/repos/%s/%s/commits?per_page=30", username, repoName)

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println("Request error:", err)
        return nil
    }

    req.Header.Set("Authorization", "Bearer "+os.Getenv("GITHUB_TOKEN"))
    req.Header.Set("Accept", "application/vnd.github+json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("API error:", err)
        return nil
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        // silently fail for private/inactive repos
        return nil
    }

    var commits []Commit
    json.NewDecoder(resp.Body).Decode(&commits)

    var times []time.Time
    for _, c := range commits {
        times = append(times, c.CommitData.Author.Date)
    }

    return times
}
