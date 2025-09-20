# task-manager-cli

A simple command-line interface (CLI) tool for managing tasks.  
Easily add, list, complete, and remove tasks directly from your terminal.

## Features

- Add new tasks
- List all tasks or filter by status (todo, in-progress, done)
- Mark tasks as completed 
- Remove tasks

## Technologies Used
- Go programming language
- Cobra for CLI framework
- JSON for data storage

## Installation

Clone the repository and install dependencies:

```bash
git clone https://github.com/chas42/task-manager-cli.git
cd task-manager-cli
go mod tidy
go build -o task-manager ./cmd
```

## Usage

```bash
./task-manager --help
./task-manager create [description]
./task-manager list
./task-manager list [status] 
./task-manager update [id] "New task description" 
./task-manager mark [status] [id]
./task-manager delete [id]
```

## License

MIT