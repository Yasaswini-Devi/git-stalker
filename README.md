# 🕵️‍♀️ git-stalker — Know a Developer, the Git Way

> A sleek, open-source CLI tool that analyzes any GitHub user's developer behavior — language preferences, commit timings, dev archetype, top repos, and more — all in one command.

---

## 📸 Sample Output


<img width="967" height="382" alt="image" src="https://github.com/user-attachments/assets/499e7bc1-6ba8-4e1c-8e10-dfd9ab01773e" />


---

## 🚀 Features

- 🔍 Profile any public GitHub user
- 💻 Analyze language usage across repositories
- ⏰ Visualize commit activity by **hour** and **day**
- 🧠 Detect developer **archetype** (Night Owl, 9-to-5 Coder, etc.)
- 🏆 List top-starred repositories
- 📝 Auto-generate a Markdown report (`--md`)
- 🖥️ Auto-open the report in your editor (`--open`)
- 🏷️ Generate a personalized badge for GitHub READMEs

---

## 📦 Installation

### 📁 Clone & Build
```bash
git clone https://github.com/YOUR_USERNAME/git-stalker.git
cd git-stalker
go build -o git-stalker
```

🔧 Or install via go install (if public)
```bash
go install github.com/YOUR_USERNAME/git-stalker@latest
```

🧪 Usage
```bash
./git-stalker <github-username> [flags]
```

🔍 Examples
```bash
./git-stalker Yasaswini-Devi
./git-stalker torvalds --md
./git-stalker gaearon --md --open
```

🛠 Flags
| Flag | Shortcut | Description |
|:-------------:|:------------:|:------------:|
| --md | -m | Generate a Markdown report |
| --open  | -o | Open the Markdown report after saving |
