package cli

import (
    "fmt"
    "log"

    "github.com/spf13/cobra"
    "github.com/artromone/4me/internal/database"
    "github.com/artromone/4me/internal/models"
)

func CreateListCommand() *cobra.Command {
    var groupID int
    var description string

    cmd := &cobra.Command{
        Use:   "create-list [name]",
        Short: "Create a new list",
        Args:  cobra.ExactArgs(1),
        Run: func(cmd *cobra.Command, args []string) {
            // Initialize database
            db := database.NewDatabase()
            defer db.Close()

            // Create list
            list := &models.List{
                Name:        args[0],
                Description: description,
                GroupID:     groupID,
            }

            id, err := db.CreateList(list)
            if err != nil {
                log.Fatalf("Failed to create list: %v", err)
            }

            fmt.Printf("List created with ID: %d\n", id)
        },
    }

    // Flags
    cmd.Flags().IntVarP(&groupID, "group", "g", 0, "ID of the group to add the list to")
    cmd.Flags().StringVarP(&description, "description", "d", "", "List description")
    cmd.MarkFlagRequired("group")

    return cmd
}
