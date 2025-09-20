package commands

import (
	"fmt"

	"github.com/chas42/task-manager-cli/internal/model"
	"github.com/chas42/task-manager-cli/internal/repository"
	"github.com/spf13/cobra"
)

func ListTask() cobra.Command {

	var command = &cobra.Command{
		Use:   "list [status]",
		Short: "List all tasks or filter by status (todo, in-progress, done)",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			var status model.Status
			if len(args) > 0 {
				status = model.Status(args[0])
			}

			if status != "" &&
				status != model.TODO &&
				status != model.DONE &&
				status != model.IN_PROGRESS {
				fmt.Println("Invalid status. Use 'todo', 'in-progress', or 'done'.")
				return
			}

			tasks, err := repository.LoadTasks(status)

			if err != nil {
				fmt.Println("Error loading tasks:", err)
				return
			}

			if len(tasks) == 0 {
				fmt.Println("No tasks found.")
				return
			}

			for _, task := range tasks {
				fmt.Printf("ID: %d, Description: %s, Status: %s\n",
					task.ID, task.Description, task.Status)
			}
		},
	}

	return *command
}
