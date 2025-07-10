package api

import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"
)

func FetchUserProfile(username string) {
    url := fmt.Sprintf("https://api.github.com/users/%s", username)

    req, _ := http.NewRequest("GET", url, nil)
    req.Header.Set("Authorization", "Bearer " + os.Getenv("GITHUB_TOKEN"))

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error fetching user:", err)
        return
    }
    defer resp.Body.Close()

    var user map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&user)

    if name, ok := user["name"].(string); ok && name != "" {
        fmt.Println("👤 Name:", name)
    } else {
        fmt.Println("👤 Name: N/A")
    }
    if bio, ok := user["bio"].(string); ok && bio != "" {
        fmt.Println("📝 Bio:", bio)
    } else {
        fmt.Println("📝 Bio: N/A")
    }
    fmt.Printf("Public Repos: %v\n", user["public_repos"])
}
