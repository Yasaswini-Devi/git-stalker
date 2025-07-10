package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/Yasaswini-Devi/git-stalker/api"
    "github.com/Yasaswini-Devi/git-stalker/analyzer"
)

var rootCmd = &cobra.Command{
    Use:   "git-stalker [username]",
    Short: "GitHub developer profiler",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        username := args[0]
        fmt.Println("Profiling:", username)
        api.FetchUserProfile(username)
        repos := api.FetchUserRepos(username)
        if len(repos) == 0 {
            fmt.Println("No repos found or error occurred.")
            return
        }
        analyzer.AnalyzeLanguages(repos)
    },
}

func Execute() {
    cobra.CheckErr(rootCmd.Execute())
}

