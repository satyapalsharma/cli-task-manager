# CLI Task Manager

A powerful and intuitive command-line interface (CLI) task manager built with Go and Cobra. Organize your daily tasks, set priorities, track due dates, and manage your to-do list directly from your terminal.

## Features

*   **File-based Persistence:** All tasks are saved to a local file, ensuring your data is never lost between sessions.
*   **Task Management:**
    *   **Add:** Create new tasks with descriptions.
    *   **List:** View all active, completed, or overdue tasks.
    *   **Complete:** Mark tasks as done.
    *   **Remove:** Delete tasks permanently.
    *   **Set/Modify:** Update task properties like description, priority, and due date.
*   **Priorities:** Assign priority levels (e.g., Low, Medium, High) to your tasks.
*   **Due Dates:** Set and track deadlines for your tasks.
*   **User-Friendly CLI:** Built with Cobra for a robust and easy-to-use command-line experience.

## Tech Stack

*   **Go:** The primary programming language for performance and concurrency.
*   **Cobra:** A powerful library for creating modern CLI applications in Go.

## Installation

### Prerequisites

*   **Go (1.18 or higher):** Ensure Go is installed and configured on your system. You can download it from [golang.org](https://golang.org/dl/).

### Build and Install

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/your-org/cli-task-manager.git
    cd cli-task-manager
    ```

2.  **Build the executable:**
    ```bash
    go build -o task .
    ```
    This will create an executable named `task` in the current directory.

3.  **Install (optional, for system-wide access):**
    To make `task` available from any directory, move it to a directory in your system's `PATH` (e.g., `/usr/local/bin` on Linux/macOS, or a custom bin directory on Windows).
    ```bash
    # On Linux/macOS
    sudo mv task /usr/local/bin/
    ```
    Now you can run `task` from anywhere.

## Usage

The `task` CLI tool provides several commands to manage your tasks.

### General Command Structure

```bash
task [command] [flags]
```

### Commands

#### `task add` - Add a new task

Adds a new task to your list.

*   **Basic:**
    ```bash
    task add "Buy groceries"
    ```
*   **With Priority:**
    Use `-p` or `--priority` flag. Valid priorities: `low`, `medium`, `high`.
    ```bash
    task add "Finish project report" -p high
    task add "Call mom" --priority medium
    ```
*   **With Due Date:**
    Use `-d` or `--due` flag. Supports various date formats (e.g., "YYYY-MM-DD", "tomorrow", "next monday").
    ```bash
    task add "Schedule dentist appointment" -d 2023-12-25
    task add "Review PR" --due tomorrow
    task add "Plan vacation" -d "next friday"
    ```
*   **Combined:**
    ```bash
    task add "Prepare presentation" -p high -d "2024-01-15"
    ```

#### `task list` - List tasks

Displays your tasks.

*   **All active tasks (default):**
    ```bash
    task list
    ```
*   **All tasks (including completed):**
    Use `-a` or `--all` flag.
    ```bash
    task list -a
    ```
*   **Completed tasks only:**
    Use `-c` or `--completed` flag.
    ```bash
    task list -c
    ```
*   **Overdue tasks only:**
    Use `-o` or `--overdue` flag.
    ```bash
    task list -o
    ```
*   **Filter by priority:**
    Use `-p` or `--priority` flag.
    ```bash
    task list -p high
    ```
*   **Sort by (e.g., priority, due date):**
    Use `-s` or `--sort` flag.
    ```bash
    task list --sort priority
    task list --sort due
    ```

#### `task do` - Mark a task as complete

Marks one or more tasks as completed using their IDs. Task IDs are shown in `task list`.

```bash
task do 1
task do 3 5 7
```

#### `task rm` - Remove a task

Removes one or more tasks using their IDs.

```bash
task rm 2
task rm 4 6
```

#### `task set` - Modify task properties

Updates the description, priority, or due date of an existing task.

*   **Change description:**
    ```bash
    task set 1 "Buy organic groceries"
    ```
*   **Change priority:**
    ```bash
    task set 1 --priority medium
    ```
*   **Change due date:**
    ```bash
    task set 1 --due "next monday"
    ```
*   **Clear due date:**
    ```bash
    task set 1 --due ""
    ```
*   **Combined:**
    ```bash
    task set 1 "Refactor authentication module" -p high -d "2024-02-01"
    ```

#### `task help` - Get help for commands

Displays help information for the main command or a specific subcommand.

```bash
task help
task help add
```

## Persistence

Tasks are stored in a JSON file located in your user's home directory.
*   **Linux/macOS:** `~/.taskmanager/tasks.json`
*   **Windows:** `C:\Users\<YourUser>\.taskmanager\tasks.json` (or similar, depending on `USERPROFILE` and `HOMEDRIVE`/`HOMEPATH`)

This ensures your tasks are saved and loaded automatically across sessions.

## Contributing

We welcome contributions! If you have suggestions for improvements, bug reports, or want to add new features, please feel free to:

1.  Fork the repository.
2.  Create a new branch (`git checkout -b feature/your-feature`).
3.  Make your changes.
4.  Commit your changes (`git commit -am 'feat: Add new feature X'`).
5.  Push to the branch (`git push origin feature/your-feature`).
6.  Open a Pull Request.

Please ensure your code adheres to the project's coding style and includes appropriate tests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.