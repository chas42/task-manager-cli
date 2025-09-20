package commands

import (
	"fmt"

	"github.com/chas42/task-manager-cli/internal/repository"
	"github.com/spf13/cobra"
)

func ListTask() cobra.Command {

	var command = &cobra.Command{
		Use:   "list [status]",
		Short: "List all tasks",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			var status repository.Status
			if len(args) > 0 {
				status = repository.Status(args[0])
			}

			if status != "" &&
				status != repository.TODO &&
				status != repository.DONE &&
				status != repository.IN_PROGRESS {
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
				fmt.Printf("ID: %d, Name: %s, Description: %s, Status: %s\n",
					task.ID, task.Name, task.Description, task.Status)
			}
		},
	}

	return *command
}
