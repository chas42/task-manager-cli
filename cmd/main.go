package main

import (
	"github.com/chas42/task-manager-cli/internal/commands"
	"github.com/spf13/cobra"
)

func main() {
	var rootCommand = &cobra.Command{}

	createTask := commands.CreateTask()
	listTask := commands.ListTask()

	rootCommand.AddCommand(&createTask)
	rootCommand.AddCommand(&listTask)
	rootCommand.Execute()
}
