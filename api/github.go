package api

import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"
)

func FetchUserProfile(username string) (string, string) {
    url := fmt.Sprintf("https://api.github.com/users/%s", username)

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println("❌ Request creation failed:", err)
        return "N/A", "N/A"
    }

    req.Header.Set("Authorization", "Bearer "+os.Getenv("GITHUB_TOKEN"))
    req.Header.Set("Accept", "application/vnd.github+json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil || resp.StatusCode != 200 {
        fmt.Println("❌ Failed to fetch user profile:", err)
        return "N/A", "N/A"
    }
    defer resp.Body.Close()

    var user map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&user)

    name := "N/A"
    bio := "N/A"

    if n, ok := user["name"].(string); ok && n != "" {
        name = n
    }
    if b, ok := user["bio"].(string); ok && b != "" {
        bio = b
    }

    fmt.Println("👤 Name:", name)
    fmt.Println("📝 Bio:", bio)
    fmt.Printf("Public Repos: %v\n", user["public_repos"])

    return name, bio
}