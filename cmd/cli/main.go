package main

import (
    "log"
    "os"

    "github.com/spf13/cobra"
    "github.com/artromone/4me/internal/cli"
)

func main() {
    rootCmd := &cobra.Command{
        Use:   "task",
        Short: "Task Management CLI",
    }

    // Add subcommands
    rootCmd.AddCommand(
        cli.CreateTaskCommand(),
        cli.ListTasksCommand(),
        cli.CreateListCommand(),
        cli.CreateGroupCommand(),
    )

    if err := rootCmd.Execute(); err != nil {
        log.Println(err)
        os.Exit(1)
    }
}
