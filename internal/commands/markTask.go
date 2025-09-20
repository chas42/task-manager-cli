package commands

import (
	"fmt"
	"strconv"

	"github.com/chas42/task-manager-cli/internal/model"
	"github.com/chas42/task-manager-cli/internal/repository"
	"github.com/spf13/cobra"
)

func MarkTask() cobra.Command {


	var command = &cobra.Command{
		Use:   "mark [status] [id]",
		Short: "Update an existent task by ID",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {

			status := model.Status(args[0])

			if status != "" &&
				status != model.TODO &&
				status != model.DONE &&
				status != model.IN_PROGRESS {
				fmt.Println("Invalid status. Use 'todo', 'in-progress', or 'done'.")
				return
			}

			taskId, _ := strconv.Atoi(args[1])
			if taskId == 0 {
				fmt.Println("You must supply a task ID.")
				return
			}

			err := repository.MarkTask(taskId, status)
			if err != nil {
				fmt.Println("Error marking task:", err)
				return
			}
			fmt.Println("Task Status changed successfully!")
		},
	}

	return *command
}
