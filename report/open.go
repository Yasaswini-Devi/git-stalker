package report

import (
    "log"
    "os/exec"
    "runtime"
)

func OpenMarkdownReport(username string) {
    filename := fmt.Sprintf("%s_report.md", username)

    cmd := exec.Command("code", filename) // 'code' CLI opens VS Code

    err := cmd.Start()
    if err != nil {
        log.Println("⚠️ Could not open in VS Code. Falling back to default opener...")

        // Optional fallback to system default opener
        switch runtime.GOOS {
        case "darwin":
            fallback = exec.Command("open", filename)
        case "windows":
            fallback = exec.Command("cmd", "/c", "start", filename)
        default:
            fallback = exec.Command("xdg-open", filename)
        }
    }
}
