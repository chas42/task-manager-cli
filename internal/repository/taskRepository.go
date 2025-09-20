package repository

import (
	"encoding/json"
	"io"
	"os"
	"time"
)

type Status string

const (
	TODO        Status = "todo"
	IN_PROGRESS Status = "in-progress"
	DONE        Status = "done"
)

type Task struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      Status `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func LoadTasks(status Status) ([]Task, error) {
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return []Task{}, nil
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}

	if status != "" {
		var filteredTasks []Task
		for _, task := range tasks {
			if task.Status == status {
				filteredTasks = append(filteredTasks, task)
			}
		}
		return filteredTasks, nil
	}

	return tasks, nil
}

func CreateTask(name string, description string) error {

	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	if len(data) == 0 {
		data = []byte("[]")
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return err
	}

	lastTaskID := 0
	if len(tasks) > 0 {
		lastTaskID = tasks[len(tasks)-1].ID
	}

	task := Task{
		ID:          lastTaskID + 1,
		Name:        name,
		Description: description,
		Status:      TODO,
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
	}

	tasks = append(tasks, task)
	file.Truncate(0)
	file.Seek(0, 0)

	encoder := json.NewEncoder(file)
	return encoder.Encode(tasks)
}
