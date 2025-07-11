package report

import (
    "fmt"
    "os/exec"
    "runtime"
)

func OpenMarkdownReport(username string) {
    filename := fmt.Sprintf("%s_report.md", username)

    var cmd *exec.Cmd

    switch runtime.GOOS {
    case "darwin":
        cmd = exec.Command("open", filename) // macOS
    case "windows":
        cmd = exec.Command("cmd", "/c", "start", filename) // Windows
    default:
        cmd = exec.Command("xdg-open", filename) // Linux
    }

    err := cmd.Start()
    if err != nil {
        fmt.Println("⚠️ Could not open file automatically:", err)
    }
}
