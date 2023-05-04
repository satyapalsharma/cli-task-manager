package cmd

import (
	"fmt"
	"sort"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"cli-task-manager/internal/store"
	"cli-task-manager/internal/task"
)

// listFlags holds the values for the command-line flags specific to the list command.
var listFlags struct {
	all     bool // --all, -a: List all tasks (pending and completed)
	done    bool // --done, -d: List only completed tasks
	pending bool // --pending, -p: List only pending tasks (default)

	dueToday bool // --due-today: List tasks due today
	overdue  bool // --overdue: List overdue tasks
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List tasks",
	Long: `List tasks in your task manager.

By default, 'list' shows all pending tasks.
You can use flags to filter by status (all, done, pending),
or by due date (due-today, overdue).`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Load tasks from the persistent store.
		tasks, err := store.LoadTasks()
		if err != nil {
			return fmt.Errorf("could not load tasks: %w", err)
		}

		// Apply filters based on command-line flags.
		filteredTasks := filterTasks(tasks)

		// Sort the filtered tasks according to defined criteria.
		sortTasks(filteredTasks)

		// Display the tasks to the console.
		displayTasks(filteredTasks)

		return nil
	},
}

func init() {
	// Add the list command as a subcommand to the root command.
	rootCmd.AddCommand(listCmd)

	// Define flags for the list command.
	listCmd.Flags().BoolVarP(&listFlags.all, "all", "a", false, "List all tasks (pending and completed)")
	listCmd.Flags().BoolVarP(&listFlags.done, "done", "d", false, "List only completed tasks")
	listCmd.Flags().BoolVarP(&listFlags.pending, "pending", "p", false, "List only pending tasks (default)")

	listCmd.Flags().BoolVar(&listFlags.dueToday, "due-today", false, "List tasks due today")
	listCmd.Flags().BoolVar(&listFlags.overdue, "overdue", false, "List overdue tasks (due date in past and not completed)")

	// Mark status flags as mutually exclusive to prevent conflicting filters.
	listCmd.MarkFlagsMutuallyExclusive("all", "done", "pending")
	// Mark date flags as mutually exclusive. A task cannot be both due today and overdue.
	listCmd.MarkFlagsMutuallyExclusive("