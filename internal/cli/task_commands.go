package cli

import (
    "fmt"
    "log"
    "time"

    "github.com/spf13/cobra"
    "github.com/artromone/4me/internal/database"
    "github.com/artromone/4me/internal/models"
)

func CreateTaskCommand() *cobra.Command {
    var listID int
    var dueDate string

    cmd := &cobra.Command{
        Use:   "create-task [title]",
        Short: "Create a new task",
        Args:  cobra.ExactArgs(1),
        Run: func(cmd *cobra.Command, args []string) {
            // Initialize database
            db := database.NewDatabase()
            defer db.Close()

            // Parse due date if provided
            var parsedDueDate time.Time
            if dueDate != "" {
                var err error
                parsedDueDate, err = time.Parse("2006-01-02", dueDate)
                if err != nil {
                    log.Fatalf("Invalid date format. Use YYYY-MM-DD")
                }
            }

            // Create task
            task := &models.Task{
                Title:    args[0],
                ListID:   listID,
                Status:   "pending",
                DueDate:  parsedDueDate,
            }

            id, err := db.CreateTask(task)
            if err != nil {
                log.Fatalf("Failed to create task: %v", err)
            }

            fmt.Printf("Task created with ID: %d\n", id)
        },
    }

    // Flags
    cmd.Flags().IntVarP(&listID, "list", "l", 0, "ID of the list to add the task to")
    cmd.Flags().StringVarP(&dueDate, "due", "d", "", "Due date (YYYY-MM-DD)")
    cmd.MarkFlagRequired("list")

    return cmd
}

func ListTasksCommand() *cobra.Command {
    var listID int

    cmd := &cobra.Command{
        Use:   "list-tasks",
        Short: "List tasks in a specific list",
        Run: func(cmd *cobra.Command, args []string) {
            // Initialize database
            db := database.NewDatabase()
            defer db.Close()

            // Fetch tasks
            tasks, err := db.ListTasks(listID)
            if err != nil {
                log.Fatalf("Failed to list tasks: %v", err)
            }

            // Print tasks
            if len(tasks) == 0 {
                fmt.Println("No tasks found.")
                return
            }

            fmt.Println("Tasks:")
            for _, task := range tasks {
                fmt.Printf("- [%s] %s (ID: %d)\n", 
                    task.Status, 
                    task.Title, 
                    task.ID,
                )
            }
        },
    }

    // Flags
    cmd.Flags().IntVarP(&listID, "list", "l", 0, "ID of the list to show tasks from")
    cmd.MarkFlagRequired("list")

    return cmd
}
