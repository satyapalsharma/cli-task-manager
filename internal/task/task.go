```go
package task

import (
	"fmt"
	"time"
)

// Priority represents the urgency level of a task.
type Priority int

// Constants for predefined priority levels.
const (
	PriorityNone   Priority = 0 // Default or unset priority.
	PriorityHigh   Priority = 1
	PriorityMedium Priority = 2
	PriorityLow    Priority = 3
)

// String returns the string representation of a Priority.
// This method makes Priority values human-readable when printed.
func (p Priority) String() string {
	switch p {
	case PriorityHigh:
		return "high"
	case PriorityMedium:
		return "medium"
	case PriorityLow:
		return "low"
	case PriorityNone:
		return "none"
	default:
		// Fallback for unexpected or invalid priority values.
		return fmt.Sprintf("unknown_priority_%d", p)
	}
}

// ParsePriorityFromString converts a string into a Priority enum.
// It returns the corresponding Priority and a boolean indicating if the parsing was successful.
// This is useful for parsing user input for priority settings.
func ParsePriorityFromString(s string) (Priority, bool) {
	switch s {
	case "high":
		return PriorityHigh, true
	case "medium":
		return PriorityMedium, true
	case "low":
		return PriorityLow, true
	case "none", "": // Allow "none" or an empty string to represent no priority.
		return PriorityNone, true
	default:
		return PriorityNone, false // Return false for unrecognized strings.
	}
}

// Task represents a single task in the CLI task manager.
// It includes fields for identification, description, status, and scheduling.
type Task struct {
	ID          int       `json:"id"`           // Unique identifier for the task, typically assigned by the store.
	Description string    `json:"description"`  // The main textual description of the task.
	Priority    Priority  `json:"priority"`     // The priority level of the task.
	DueDate     time.Time `json:"due_date"`     // Optional due date for the task. A zero value indicates no due date.
	CreatedAt   time.Time `json:"created_at"`   // Timestamp when the task was created.
	CompletedAt time.Time `json:"completed_at"` // Timestamp when the task was completed. A zero value indicates not completed.
}

// NewTask creates and initializes a new Task with the given description
// and default values for other fields (e.g., no priority, creation time set to now).
func NewTask(description string) *Task {
	return &Task{
		Description: description,
		Priority:    PriorityNone, // Tasks start with no specific priority.
		CreatedAt:   time.Now(),   // Set creation time to the current moment.
		// DueDate and CompletedAt are time.Time zero values by default,
		// which correctly indicates they are unset.
	}
}

// IsCompleted checks if the task has been marked as completed.
// A task is considered completed if its CompletedAt timestamp is not the zero value.
func (t *Task) IsCompleted() bool {
	return !t.CompletedAt.IsZero()
}

// IsOverdue checks if the task has a due date that has passed and is not yet completed.
// A task is overdue if its DueDate is set, is before the current time, and the task is not completed.
func (t *Task) IsOverdue() bool {
	if t.DueDate.IsZero() || t.IsCompleted() {
		return false // Not overdue if there's no due date or if the task is already completed.
	}
	// Compare the due date with the current time.
	return t.DueDate.Before(time.Now())
}
```