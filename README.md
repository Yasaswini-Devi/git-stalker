# 🕵️‍♀️ Git Stalker

> 🔖 **Phase 1: Completed**
> ✅ GitHub CLI profiler with markdown reports and archetype classification.
> 🚧 Phase 2: Planned — AI enhancements, GraphQL support, and richer contribution analysis.

![GitHub repo size](https://img.shields.io/github/repo-size/Yasaswini-Devi/git-stalker)
![Last Commit](https://img.shields.io/github/last-commit/Yasaswini-Devi/git-stalker)
![License](https://img.shields.io/github/license/Yasaswini-Devi/git-stalker)
![Dev Archetype](https://img.shields.io/badge/Archetype-9_to_5_Coder-blueviolet)

Git Stalker is a developer profiler CLI tool built with Go that analyzes a GitHub user's public activity and summarizes:

- 💻 Language usage across public repos
- ⏰ Commit activity by hour and day
- 🧠 Developer archetype classification
- 🏆 Top-starred public repositories
- 📄 Optionally generates a Markdown report with badge

---

## 🚀 Installation

```bash
go install github.com/Yasaswini-Devi/git-stalker@latest
```

Or clone and run:

```bash
git clone https://github.com/Yasaswini-Devi/git-stalker
cd git-stalker
go run main.go [username]
```

> ⚠️ You must set your GitHub token in the environment:
```bash
export GITHUB_TOKEN=your_personal_access_token
```

---

## 📦 Usage

```bash
go run main.go [username] [flags]
```

### Example:

```bash
go run main.go Yasaswini-Devi --md --open --log
```

### Flags:

| Flag | Description |
|------|-------------|
| `--md`, `-m` | Generate a markdown report |
| `--open`, `-o` | Open the report in VS Code |
| `--log`, `-l` | Enable verbose logging |

---

## 📝 Sample Output

```
👤 Name: Yasaswini Devi
📝 Bio: Builder of things, breaker of silence
Public Repos: 9

💻 Language Usage:
• Python          : 2 repos
• Go              : 1 repos
• Java            : 1 repos
• JavaScript      : 3 repos

⏰ Commit Activity:
- Peak hour: 17:00 → 11 commits
- Most active day: Friday → 26 commits

🏆 Top Starred Repositories:
• ATM-Simulator – ⭐ 0 stars
• Book-Recommendation – ⭐ 0 stars
```

---

## 📄 Markdown Report

If `--md` is enabled, a Markdown file like `Yasaswini-Devi_report.md` is generated with all analysis and a badge you can add to your README:

```markdown
![Dev Archetype](https://img.shields.io/badge/Archetype-9_to_5_Coder-blueviolet)
```

---

## 🛠️ Tech Stack

- [Go](https://golang.org/)
- [Cobra](https://github.com/spf13/cobra) for CLI
- GitHub REST API v3 (GraphQL in future)
- Optional: VS Code CLI for preview

---

## 📜 License

This project is licensed under the [MIT License](LICENSE).

---

## 🙌 Contributions

Feel free to fork, raise issues, or suggest new features like:
- GraphQL API integration
- Contributor heatmaps
- AI-powered personality predictions

---

## 🌟 Star the repo if you like it!