```go
package cmd

import (
	"fmt"
	"os"
	"strconv"
	"time" // Required for setting task completion time

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/cli-task-manager/internal/store" // Adjust this import path to your project's module path
	"github.com/cli-task-manager/internal/task"   // Adjust this import path to your project's module path
)

// doCmd represents the 'do' command
var doCmd = &cobra.Command{
	Use:   "do [TASK_ID...]",
	Short: "Marks one or more tasks as complete",
	Long: `Marks one or more tasks as complete by their ID.
The task ID refers to the number displayed when listing tasks (e.g., using 'task list').

Example:
  task do 1       - Marks the first task in the list as complete.
  task do 2 4     - Marks the second and fourth tasks as complete.`,
	Args: cobra.MinimumNArgs(1), // The 'do' command requires at least one task ID argument.
	RunE: func(cmd *cobra.Command, args []string) error {
		// Retrieve the data file path from Viper configuration.
		// Viper reads from a config file (e.g., ~/.config/task/config.yaml) or environment variables.
		dataFilePath := viper.GetString("datafile")
		if dataFilePath == "" {
			// If the data file path is not configured, return an error.
			return fmt.Errorf("data file path not set. Please configure it using 'task set datafile <path>' or ensure it's in your config file")
		}

		// Initialize the task store with the determined data file path.
		s := store.NewStore(dataFilePath)

		// Load existing tasks from the store.
		tasks, err := s.LoadTasks()
		if err != nil {
			// If the data file does not exist, it means there are no tasks yet.
			if os.IsNotExist(err) {
				return fmt.Errorf("no tasks found. The data file '%s' does not exist or is empty. Use 'task add' to create tasks", dataFilePath)
			}
			// Handle other potential errors during task loading.
			return fmt.Errorf("failed to load tasks: %w", err)
		}

		// If no tasks are loaded (e.g., file exists but is empty), inform the user.
		if len(tasks) == 0 {
			return fmt.Errorf("no tasks available to mark as complete. Use 'task add' to create tasks")
		}

		var tasksCompletedCount int // Counter for successfully marked tasks
		var tasksProcessedCount int // Counter for task IDs attempted to process

		// Iterate through each argument, which is expected to be a task ID.
		for _, arg := range args {
			tasksProcessedCount++ // Increment counter for each argument processed
			id, err := strconv.Atoi(arg)
			if err != nil {
				// If the argument is not a valid integer, print an error and continue to the next argument.
				fmt.Printf("Error: Invalid task ID '%s'. Please provide a number.\n", arg)
				continue
			}

			// Convert the 1-based user ID (