package cmd

import (
	"fmt"
	"os"
	"path/filepath" // For cross-platform path manipulation

	"github.com/spf13/cobra"
)

var (
	// dataFile stores the path to the JSON file where tasks are persisted.
	// This variable will be set by a persistent flag or a default value.
	// It is accessible by all subcommands to initialize the task store.
	dataFile string
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "A CLI task manager",
	Long: `task is a simple yet powerful command-line task manager.
It allows you to add, list, complete, remove, and modify tasks
with file-based persistence, priorities, and due dates.

It stores tasks in a JSON file, by default located in your user's
configuration directory (e.g., ~/.taskmanager/tasks.json on Linux/macOS).
You can specify an alternative data file using the --data-file flag.`,
	// For a root command that primarily dispatches to subcommands,
	// a direct Run function is often not needed unless it provides a
	// default action (e.g., showing help or listing tasks if no subcommand is given).
	// Run: func(cmd *cobra.Command, args []string) {
	//     // Example: if no subcommand is given, show help
	//     cmd.Help()
	// },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is the main entry point for the Cobra application, called by main.main().
// It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		// Cobra handles most errors by printing them, but for critical
		// errors or custom error messages, we can explicitly print and exit.
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// init is called before main.main() and is used to initialize
// flags and add subcommands to the root command.
func init() {
	// Define a persistent flag for the data file path.
	// This flag will be available to all subcommands in this application.
	// It allows users to specify an alternative location for task data.
	defaultDataFile := getDefaultDataFilePath()
	rootCmd.PersistentFlags().StringVarP(&dataFile, "data-file", "d", defaultDataFile, "Path to the JSON file for task storage")

	// Add all subcommands to the root command.
	// These commands are defined in their respective files (e.g., cmd/add.go).
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(doCmd)
	rootCmd.AddCommand(rmCmd)
	rootCmd.AddCommand(setCmd)

	// Cobra also supports local flags, which will only run when this command
	// is called directly, not by its subcommands.
	// Example: rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// getDefaultDataFilePath determines the default path for the task data file.
// It attempts to place the file in a user-specific configuration directory
// (e.g., ~/.taskmanager/tasks.json on Linux/macOS, or equivalent on Windows).
// If the directory doesn't exist, it attempts to create it.
// If any errors occur, it falls back to "tasks.json" in the current working directory.
func getDefaultDataFilePath() string {
	// Attempt to get the user's home directory.
	home, err := os.UserHomeDir()
	if err != nil {
		// Fallback to current directory if home dir cannot be determined.
		fmt.Fprintf(os.Stderr, "Warning: Could not determine user home directory (%v). Using 'tasks.json' in current