package api

import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"
)

type Repo struct {
    Name            string `json:"name"`
    Language        string `json:"language"`
    StargazersCount int    `json:"stargazers_count"`
}

func FetchUserRepos(username string) []Repo {
    url := fmt.Sprintf("https://api.github.com/users/%s/repos?per_page=100", username)

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
        fmt.Println("API call error:", err)
        return nil
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        fmt.Println("Error: GitHub API returned", resp.Status)
        return nil
    }

    var repos []Repo
    err = json.NewDecoder(resp.Body).Decode(&repos)
    if err != nil {
        fmt.Println("JSON decode error:", err)
        return nil
    }

    return repos
}
