package cmd

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/your-username/cli-task-manager/internal/store"
	"github.com/your-username/cli-task-manager/internal/task"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <description>",
	Short: "Add a new task to your list",
	Long: `Add a new task with a given description.
You can optionally specify a priority and a due date.

Examples:
  taskman add "Buy groceries"
  taskman add "Finish report" -p 2
  taskman add "Call mom" -d 2023-12-25
  taskman add "Plan vacation" -p 1 -d 2024-07-15`,
	Args: cobra.MinimumNArgs(1), // Requires at least one argument for the task description
	RunE: func(cmd *cobra.Command, args []string) error {
		// Initialize the task store.
		// In a real application, the store might be initialized once in root.go
		// and passed down or made globally accessible. For simplicity in this file,
		// we initialize it here.
		s, err := store.NewStore()
		if err != nil {
			return fmt.Errorf("failed to initialize task store: %w", err)
		}

		// Extract task description from arguments.
		description := args[0]

		// Create a new task object.
		newTask := task.Task{
			Description: description,
			Done:        false, // New tasks are always not done
			CreatedAt:   time.Now(),
		}

		// Handle optional priority flag.
		priorityStr, _ := cmd.Flags().GetString("priority")
		if priorityStr != "" {
			p, err := strconv.Atoi(priorityStr)
			if err != nil {
				return fmt.Errorf("invalid priority value '%s': must be an integer", priorityStr)
			}
			if p < 0 {
				return fmt.Errorf("priority cannot be negative")
			}
			newTask.Priority = task.Priority(p)
		}

		// Handle optional due date flag.
		dueDateStr, _ := cmd.Flags().GetString("due")
		if dueDateStr != "" {
			// Define the expected date format (YYYY-MM-DD).
			const dateFormat = "2006-01-02"
			dueDate, err := time.Parse(dateFormat, dueDateStr)
			if err != nil {
				return fmt.Errorf("invalid due date format '%s': expected YYYY-MM-DD", dueDateStr)
			}
			newTask.DueDate = &dueDate // Assign the parsed time to the pointer
		}

		// Add the new task to the store.
		if err := s.AddTask(newTask); err != nil {
			return fmt.Errorf("failed to add task: %w", err)
		}

		fmt.Printf("Task added: \"%s\"\n", newTask.Description)
		if newTask.Priority > 0 {
			fmt.Printf("Priority: %d\n", newTask.Priority)
		}
		if newTask.DueDate != nil {
			fmt.Printf("Due: %s\n", newTask.DueDate.Format("2006-01-02"))
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Define flags for the add command.
	// --priority or -p: Allows setting a task's priority (integer).
	addCmd.Flags().StringP("priority", "p", "", "Set the priority of the task (e.g., 1 for high, 0 for no priority)")

	// --due or -d: Allows setting a task's due date (string in YYYY-MM-DD format).
	addCmd.Flags().StringP("due", "d", "", "Set the due date of the task (format: YYYY-MM-DD)")
}