package commands

import (
	"fmt"

	"github.com/chas42/task-manager-cli/internal/repository"
	"github.com/spf13/cobra"
)

func CreateTask() cobra.Command {

	var command = &cobra.Command{
		Use:   "create [description]",
		Short: "Create a new task",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			taskDescription := args[0]
			if taskDescription == "" {
				fmt.Println("You must supply a task description.")
				return
			}

			err := repository.CreateTask(taskDescription)
			if err != nil {
				fmt.Println("Error creating task:", err)
				return
			}
			fmt.Println("Task created successfully!")
		},
	}

	return *command
}
