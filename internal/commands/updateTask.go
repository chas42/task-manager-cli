package commands

import (
	"fmt"
	"strconv"

	"github.com/chas42/task-manager-cli/internal/repository"
	"github.com/spf13/cobra"
)

func UpdateTask() cobra.Command {

	var command = &cobra.Command{
		Use:   "update [id] [description]",
		Short: "Update an existent task by ID",
		Run: func(cmd *cobra.Command, args []string) {

			taskId := 0
			taskId, _ = strconv.Atoi(args[0])
			if taskId == 0 {
				fmt.Println("You must supply a task ID.")
				return
			}

			taskDescription := args[1]
			if taskDescription == "" {
				fmt.Println("You must supply a new task description.")
				return
			}

			err := repository.UpdateTask(taskId, taskDescription)
			if err != nil {
				fmt.Println("Error updating task:", err)
				return
			}
			fmt.Println("Task updated successfully!")
		},
	}

	return *command
}
