package commands

import (
	"fmt"
	"strconv"

	"github.com/chas42/task-manager-cli/internal/repository"
	"github.com/spf13/cobra"
)

func DeleteTask() cobra.Command {

	var command = &cobra.Command{
		Use:   "delete [id]",
		Short: "Delete a task by ID",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			taskId := args[0]
			taskIdInt, err := strconv.Atoi(taskId)
			if err != nil {
				fmt.Println("Invalid task ID:", err)
				return
			}

			err = repository.DeleteTask(taskIdInt)
			if err != nil {
				fmt.Println("Error deleting task:", err)
				return
			}

			fmt.Println("Task deleted successfully!")
		},
	}

	return *command
}
