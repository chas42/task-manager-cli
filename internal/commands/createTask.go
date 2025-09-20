package commands

import (
	"fmt"

	"github.com/chas42/task-manager-cli/internal/repository"
	"github.com/spf13/cobra"
)

func CreateTask() cobra.Command {

	var taskName, taskDescription string

	var command = &cobra.Command{
		Use:   "create",
		Short: "Create a new task",
		Run: func(cmd *cobra.Command, args []string) {
			// validations
			if taskName == "" {
				fmt.Println("You must supply a task name.")
				return
			}
			if taskDescription == "" {
				fmt.Println("You must supply a task description.")
				return
			}

			err := repository.CreateTask(taskName, taskDescription)
			if err != nil {
				fmt.Println("Error creating task:", err)
				return
			}
			fmt.Println("Task created successfully!")
		},
	}

	command.Flags().StringVarP(&taskName, "name", "n", "", "Name of the task")
	command.Flags().StringVarP(&taskDescription, "description", "d", "", "Description of the task")

	return *command
}
