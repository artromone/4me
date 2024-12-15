package cli

import (
    "fmt"
    "log"

    "github.com/spf13/cobra"
    "github.com/artromone/4me/internal/database"
    "github.com/artromone/4me/internal/models"
)

func CreateGroupCommand() *cobra.Command {
    var description string

    cmd := &cobra.Command{
        Use:   "create-group [name]",
        Short: "Create a new group",
        Args:  cobra.ExactArgs(1),
        Run: func(cmd *cobra.Command, args []string) {
            // Initialize database
            db := database.NewDatabase()
            defer db.Close()

            // Create group
            group := &models.Group{
                Name:        args[0],
                Description: description,
            }

            id, err := db.CreateGroup(group)
            if err != nil {
                log.Fatalf("Failed to create group: %v", err)
            }

            fmt.Printf("Group created with ID: %d\n", id)
        },
    }

    // Flags
    cmd.Flags().StringVarP(&description, "description", "d", "", "Group description")

    return cmd
}

func ListGroupsCommand() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "list-groups",
        Short: "List all groups",
        Run: func(cmd *cobra.Command, args []string) {
            // Initialize database
            db := database.NewDatabase()
            defer db.Close()

            // Fetch groups
            groups, err := db.ListGroups()
            if err != nil {
                log.Fatalf("Failed to list groups: %v", err)
            }

            // Print groups
            if len(groups) == 0 {
                fmt.Println("No groups found.")
                return
            }

            fmt.Println("Groups:")
            for _, group := range groups {
                fmt.Printf("- %s (ID: %d)\n", group.Name, group.ID)
                if group.Description != "" {
                    fmt.Printf("  Description: %s\n", group.Description)
                }
            }
        },
    }

    return cmd
}
