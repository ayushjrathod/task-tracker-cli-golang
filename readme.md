# Task Manager CLI

A lightweight, efficient command-line task management tool built in Go. Manage your tasks with simple commands while maintaining data persistence through local JSON storage.

[![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

## Features

- ✅ **Add Tasks**: Create new tasks with descriptive names
- 📋 **List Tasks**: View all tasks with their completion status
- ❌ **Remove Tasks**: Delete unwanted tasks by ID
- ✔️ **Mark Complete**: Mark tasks as done
- 🗑️ **Clear All**: Remove all tasks at once
- 💾 **Persistent Storage**: Tasks saved in local JSON file
- 🚀 **Fast & Lightweight**: Minimal dependencies, quick execution

## Installation

### Prerequisites

- Go 1.21 or higher
- Git (for cloning the repository)

### Build from Source

```bash
# Clone the repository
git clone <repository-url>
cd task-manager

# Build the application
go build -o task-manager .

# Run the application
./task-manager
```

### Alternative: Direct Go Install

```bash
go install github.com/your-username/task-manager@latest
```

## Usage

### Starting the Application

```bash
./task-manager
```

The application will start in interactive mode, displaying:

```
Welcome to Task Manager CLI
Type 'exit' to quit the program
Enter command:
```

### Available Commands

| Command  | Syntax             | Description            | Example               |
| -------- | ------------------ | ---------------------- | --------------------- |
| `add`    | `add <task_name>`  | Create a new task      | `add "Buy groceries"` |
| `list`   | `list`             | Display all tasks      | `list`                |
| `remove` | `remove <task_id>` | Delete a specific task | `remove 1`            |
| `done`   | `done <task_id>`   | Mark task as completed | `done 1`              |
| `clear`  | `clear` or `reset` | Remove all tasks       | `clear`               |
| `exit`   | `exit`             | Quit the application   | `exit`                |

### Example Session

```bash
Enter command: add "Learn Go programming"
Enter command: add "Build a CLI tool"
Enter command: add "Write documentation"
Enter command: list
ID: 1 | Learn Go programming | Status: Not Done
ID: 2 | Build a CLI tool | Status: Not Done
ID: 3 | Write documentation | Status: Not Done

Enter command: done 2
Enter command: list
ID: 1 | Learn Go programming | Status: Not Done
ID: 2 | Build a CLI tool | Status: Done
ID: 3 | Write documentation | Status: Not Done

Enter command: remove 1
Enter command: list
ID: 2 | Build a CLI tool | Status: Done
ID: 3 | Write documentation | Status: Not Done

Enter command: exit
```

## Data Storage

Tasks are automatically persisted in a `tasks.json` file located in the same directory as the executable. The file structure:

```json
[
  {
    "id": "1",
    "task_name": "Example Task",
    "status": false
  }
]
```

- **id**: Unique identifier for each task
- **task_name**: Description of the task
- **status**: Boolean indicating completion (true = done, false = not done)

## Project Structure

```
task-manager/
├── main.go              # Application entry point and CLI interface
├── go.mod              # Go module definition
├── tasks.json          # Data storage (created automatically)
├── task/
│   └── task.go         # Task business logic
├── storage/
│   └── storage.go      # File I/O operations
└── utils/
    └── utils.go        # Utility functions
```

## Error Handling

The application handles common error scenarios gracefully:

- **Missing task name**: Prompts for proper usage when adding tasks
- **Invalid task ID**: Shows error for non-existent or malformed IDs
- **Empty task list**: Displays helpful message when no tasks exist
- **File system errors**: Handles file creation and access issues

## Development

### Running Tests

```bash
go test ./...
```

### Building for Different Platforms

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o task-manager-linux .

# macOS
GOOS=darwin GOARCH=amd64 go build -o task-manager-macos .

# Windows
GOOS=windows GOARCH=amd64 go build -o task-manager-windows.exe .
```

### Code Style

This project follows standard Go conventions:

- `gofmt` for formatting
- `go vet` for static analysis
- Descriptive variable and function names
- Proper error handling

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Roadmap

- [ ] Task priorities and categories
- [ ] Due dates and reminders
- [ ] Task search and filtering
- [ ] Export/import functionality
- [ ] Configuration file support
- [ ] Colored output for better UX

## Support

If you encounter any issues or have questions:

1. Check the [Issues](../../issues) page
2. Create a new issue with detailed information
3. Include your Go version and operating system

---

**Made with ❤️ in Go**
