```go
package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"taskmanager/internal/store"
	"taskmanager/internal/task"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set <id> [flags]",
	Short: "Modify properties of an existing task",
	Long: `Use 'set' to change properties like priority or due date of an existing task.
You must provide the task ID.

Examples:
  task set 1 --priority 1
  task set 2 --due "2023-12-31"
  task set 3 -p 2 -d "tomorrow"`,
	Args: cobra.ExactArgs(1), // Requires exactly one argument: the task ID
	RunE: setRun,
}

func init() {
	rootCmd.AddCommand(setCmd)

	// Define flags for the set command
	setCmd.Flags().IntP("priority", "p", 0, "Set the priority of the task (1-3, 1 being highest)")
	setCmd.Flags().StringP("due", "d", "", "Set the due date of the task (YYYY-MM-DD, 'today', 'tomorrow', 'yesterday')")
}

// setRun is the main function executed when the 'set' command is called.
func setRun(cmd *cobra.Command, args []string) error {
	// 1. Parse Task ID from arguments
	taskIDStr := args[0]
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		return fmt.Errorf("invalid task ID: %s. Please provide a number", taskIDStr)
	}

	// 2. Initialize the task store
	dataFilePath := viper.GetString("datafile")
	s := store.NewStore(dataFilePath)

	// 3. Load existing tasks
	tasks, err := s.LoadTasks()
	if err != nil {
		return fmt.Errorf("failed to load tasks: %w", err)
	}

	// 4. Validate Task ID and find the task
	if taskID <= 0 || taskID > len(tasks) {
		return fmt.Errorf("task with ID %d not found. Please provide a valid task ID", taskID)
	}
	taskIndex := taskID - 1 // Convert 1-based CLI ID to 0-based slice index
	t := &tasks[taskIndex]  // Get a pointer to the task to modify it directly in the slice

	// Check if any flags were provided
	prioritySet := cmd.Flags().Changed("priority")
	dueSet := cmd.Flags().Changed("due")

	if !prioritySet && !dueSet {
		fmt.Println("No properties to set. Use --priority or --due flags.")
		return nil // Not an error, just nothing to do
	}

	updated := false // Flag to track if any changes were actually made

	// 5. Process --priority flag if set
	if prioritySet {
		priority, err := cmd.Flags().GetInt("priority")
		if err != nil {
			return fmt.Errorf("failed to parse priority: %w", err)
		}
		if priority < 1 || priority > 3 {
			return fmt.Errorf("invalid priority value: %d. Priority must be between 1 (highest) and 3 (lowest)", priority)
		}
		if t.Priority != priority {
			t.Priority = priority
			updated = true
			fmt.Printf("Task %d priority set to %d.\n", taskID, priority)
		} else {
			fmt.Printf("Task %d priority is already %d. No change made.\n", taskID, priority)
		}
	}

	// 6. Process --due flag if set
	if dueSet {
		dueDateStr, err := cmd.Flags().GetString("due")
		if err != nil {
			return fmt.Errorf("failed to parse due date string: %w", err)
		}

		// Handle clearing the due date if a specific keyword is used, e.g., "none"
		if strings.ToLower(strings.TrimSpace(dueDateStr)) == "none" {
			if t.DueDate != nil {
				t.DueDate = nil // Set DueDate to nil to clear it
				updated = true
				fmt.Printf("Task %d due date cleared.\n", taskID)
			} else {
				fmt.Printf("Task %d already has no due date. No change made.\n", taskID)
			}
		} else {
			dueDate, err := parseDueDate(dueDateStr)
			if err != nil {
				return fmt.Errorf("invalid due date format for '%s': %w", dueDateStr, err)
			}

			// Compare existing due date with new one to avoid unnecessary updates
			if t.DueDate == nil || !t.DueDate.Equal(dueDate) {
				t.DueDate = &dueDate
				updated = true
				fmt.Printf("Task %d due date set to %s.\n", taskID, dueDate.Format("2006-01-02"))
			} else {
				fmt.Printf("Task %d due date is already %s. No change made.\n", taskID, dueDate.Format("2006-01-02"))
			}
		}
	}

	// 7. Save tasks back to the store if any changes were made
	if updated {
		err = s.SaveTasks(tasks)
		if err != nil {
			return fmt.Errorf("failed to save tasks: %w", err)
		}
		fmt.Printf("Task %d updated successfully.\n", taskID)
	} else {
		fmt.Println("No changes were made to the task.")
	}

	return nil
}

// parseDueDate