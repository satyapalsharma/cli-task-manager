// Package store provides functionality for persisting and retrieving tasks from a file.
package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/your-username/cli-task-manager/internal/task" // Adjust import path as necessary
)

// Store manages the persistence of tasks to and from a file.
type Store struct {
	filePath string
}

// NewStore creates and initializes a new Store instance.
// It takes the path to the data file where tasks will be stored.
func NewStore(filePath string) *Store {
	return &Store{
		filePath: filePath,
	}
}

// LoadTasks reads tasks from the configured file path.
// If the file does not exist or is empty, it returns an empty slice of tasks
// without an error, indicating no tasks are currently stored.
// It returns an error if there's an issue reading the file or unmarshaling its content.
func (s *Store) LoadTasks() ([]task.Task, error) {
	// Check if the file exists. If not, return an empty list.
	if _, err := os.Stat(s.filePath); os.IsNotExist(err) {
		return []task.Task{}, nil // No file, no tasks, not an error.
	}

	// Read the file content.
	data, err := os.ReadFile(s.filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read task file: %w", err)
	}

	// If the file is empty, return an empty list.
	if len(data) == 0 {
		return []task.Task{}, nil
	}

	var tasks []task.Task
	// Unmarshal the JSON data into a slice of Task structs.
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, fmt.Errorf("failed to unmarshal tasks from file: %w", err)
	}

	return tasks, nil
}

// SaveTasks writes the given slice of tasks to the configured file path.
// It marshals the tasks into JSON format and writes them to the file,
// overwriting any existing content.
// It ensures the directory for the file exists before writing.
// It returns an error if there's an issue marshaling tasks or writing to the file.
func (s *Store) SaveTasks(tasks []task.Task) error {
	// Marshal the tasks slice into JSON format.
	// json.MarshalIndent is used for pretty-printing the JSON, making the file human-readable.
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal tasks to JSON: %w", err)
	}

	// Ensure the directory for the file exists.
	dir := filepath.Dir(s.filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %q: %w", dir, err)
	}

	// Write the JSON data to the file.
	// os.WriteFile creates the file if it doesn't exist, or truncates it if it does.
	// The file permission 0644 means owner can read/write, others can read.
	if err := os.WriteFile(s.filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write tasks to file: %w", err)
	}

	return nil
}

// GetNextID calculates the next available ID for a new task.
// It iterates through the existing tasks and returns the maximum ID found + 1.
// If no tasks exist, it returns 1.
func (s *Store) GetNextID(tasks []task.Task) int {
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	return maxID + 1
}

// FindTaskByID finds a task by its ID in the given slice of tasks.
// It returns the task and its index if found, otherwise an error.
func (s *Store) FindTaskByID(tasks []task.Task, id int) (task.Task, int, error) {
	for i, t := range tasks {
		if t.ID == id {
			return t, i, nil
		}
	}
	return task.Task{}, -1, errors.New("task not found")
}