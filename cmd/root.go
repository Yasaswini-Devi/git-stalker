package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/Yasaswini-Devi/git-stalker/api"
)

var rootCmd = &cobra.Command{
    Use:   "git-stalker [username]",
    Short: "GitHub developer profiler",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        username := args[0]
        fmt.Println("Profiling:", username)
        api.FetchUserProfile(username)
    },
}

func Execute() {
    cobra.CheckErr(rootCmd.Execute())
}

