package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/your-username/task-manager/internal/store" // Adjust this import path to your project's structure
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm [task ID...]",
	Short: "Removes one or more tasks by their ID",
	Long: `Removes one or more tasks from your task list.
You must provide at least one task ID.
Example:
  task rm 1
  task rm 1 3 5`,
	Args: cobra.MinimumNArgs(1), // Ensures at least one argument (task ID) is provided
	RunE: func(cmd *cobra.Command, args []string) error {
		// Retrieve the data file path from Viper configuration.
		// This path is typically set up in the root command's initialization.
		storePath := viper.GetString("datafile")
		if storePath == "" {
			return fmt.Errorf("data file path not configured. Please ensure 'datafile' is set in your configuration or via a flag")
		}

		// Initialize the task store with the determined path.
		// This store handles the persistence of tasks.
		s, err := store.NewStore(storePath)
		if err != nil {
			return fmt.Errorf("failed to initialize task store: %w", err)
		}
		defer s.Close() // Ensure the store's resources (e.g., file handles) are properly closed.

		var removedIDs []int          // To keep track of successfully removed task IDs
		var failedRemovals []string   // To collect error messages for tasks that couldn't be removed

		// Iterate through all provided arguments, treating each as a potential task ID.
		for _, arg := range args {
			id, err := strconv.Atoi(arg) // Convert the string argument to an integer ID.
			if err != nil {
				// If conversion fails, record the invalid argument and continue to the next.
				failedRemovals = append(failedRemovals, fmt.Sprintf("'%s' (not a valid number)", arg))
				continue
			}

			// Attempt to remove the task from the store.
			err = s.RemoveTask(id)
			if err != nil {
				// If removal fails (e.g., task not found, or other store error),
				// record the ID and the specific error message.
				failedRemovals = append(failedRemovals, fmt.Sprintf("%d (%v)", id, err))
			} else {
				// If successful, add the ID to our list of removed tasks.
				removedIDs = append(removedIDs, id)
			}
		}

		// Provide feedback to the user based on the processing results.
		if len(removedIDs) > 0 {
			fmt.Printf("Successfully removed tasks: %s\n", formatIDs(removedIDs))
		}

		if len(failedRemovals) > 0 {
			// If there were any failures, report them and return an error to indicate partial or full failure.
			return fmt.Errorf("failed to remove some tasks: %s", strings.Join(failedRemovals, ", "))
		}

		return nil // All specified tasks were processed successfully, or no critical errors occurred.
	},
}

func init() {
	// Add the 'rmCmd' to the root command, making it available as 'task rm'.
	rootCmd.AddCommand(rmCmd)
}

// formatIDs is a helper function to format a slice of integers into a comma-separated string.
// This improves the readability of output messages.
func formatIDs(ids []int) string {
	if len(ids) == 0 {
		return ""
	}
	strIDs := make([]string, len(ids))
	for i, id := range ids {
		strIDs[i] = strconv.Itoa(id)
	}
	return strings.Join(strIDs, ", ")
}