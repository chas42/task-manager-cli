# task-manager-cli

A simple command-line interface (CLI) tool for managing tasks.  
Easily add, list, complete, and remove tasks directly from your terminal.

## Features

- Add new tasks
- List all tasks or filter by status (todo, in-progress, done)
- Mark tasks as completed (not implemented yet)
- Remove tasks (not implemented yet)

## Installation

Clone the repository and install dependencies:

```bash
git clone https://github.com/chas42/task-manager-cli.git
cd task-manager-cli
go mod tidy
go build -o task-manager
```

## Usage

```bash
./task-manager create "Your task description"
./task-manager list
./task-manager list todo
./task-manager list in-progress
./task-manager list done
```

## License

MIT