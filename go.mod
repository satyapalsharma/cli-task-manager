module github.com/yourorg/cli-task-manager

// Specify the Go language version required for this module.
// It's good practice to use a recent stable version.
go 1.22

// Define the external dependencies required by this project.
// Each dependency is listed with its module path and a specific version.
require (
	// Cobra is a library for creating powerful modern CLI applications.
	// It provides a framework for defining commands, arguments, and flags.
	github.com/spf13/cobra v1.8.0

	// Viper is a complete configuration solution for Go applications.
	// It's often used alongside Cobra to handle configuration files,
	// environment variables, and command-line flags seamlessly.
	github.com/spf13/viper v1.18.2
)

// The 'go.mod' file might also contain 'replace' or 'exclude' directives
// for specific development or dependency management scenarios, but they
// are not typically needed for a standard project setup.
//
// 'replace' directives are used to substitute a module path with a different
// local path or a different remote module, often for local development
// of a dependency or forking.
//
// 'exclude' directives are used to prevent specific versions of a module
// from being used.