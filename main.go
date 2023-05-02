package main

import (
	"log"

	"github.com/taskmanager-org/cli-task-manager/cmd" // Adjust this import path to your actual module path
)

// main is the entry point of the CLI application.
// It initializes and executes the root command defined in the 'cmd' package.
func main() {
	// Execute the root command of the CLI application.
	// This call parses command-line arguments, flags, and executes the appropriate subcommand.
	// All subcommands and their logic are orchestrated through this single entry point.
	if err := cmd.Execute(); err != nil {
		// If an error occurs during command execution (e.g., invalid command, flag parsing error,
		// or an error propagated from a subcommand), log the error and terminate the program.
		// Cobra's Execute() method typically prints user-friendly errors to stderr,
		// but logging it here provides an additional record and ensures a non-zero exit status
		// for scripting purposes.
		log.Fatalf("Error executing CLI command: %v", err)
		// log.Fatalf automatically calls os.Exit(1) after printing the message.
	}
}