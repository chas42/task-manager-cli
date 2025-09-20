package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/chas42/task-manager-cli/internal/model"
)

const taskFile = "data/tasks.json"

func LoadTasksFromFile() ([]model.Task, error) {
	file, err := os.OpenFile(taskFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return []model.Task{}, nil
	}

	var tasks []model.Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func saveTasks(tasks []model.Task) error {
	file, err := os.OpenFile(taskFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	file.Truncate(0)
	file.Seek(0, 0)

	encoder := json.NewEncoder(file)
	return encoder.Encode(tasks)
}

func LoadTasks(status model.Status) ([]model.Task, error) {
	LoadTasksFromFile()
	tasks, err := LoadTasksFromFile()
	if err != nil {
		return nil, err
	}

	if status != "" {
		var filteredTasks []model.Task
		for _, task := range tasks {
			if task.Status == status {
				filteredTasks = append(filteredTasks, task)
			}
		}
		return filteredTasks, nil
	}

	return tasks, nil
}

func CreateTask(description string) error {

	tasks, err := LoadTasksFromFile()
	if err != nil {
		return err
	}

	lastTaskID := 0
	if len(tasks) > 0 {
		lastTaskID = tasks[len(tasks)-1].ID
	}

	task := model.Task{
		ID:          lastTaskID + 1,
		Description: description,
		Status:      model.TODO,
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
	}

	tasks = append(tasks, task)

	err = saveTasks(tasks)
	if err != nil {
		return err
	}

	return nil
}

func UpdateTask(taskId int, description string) error {
	tasks, err := LoadTasksFromFile()
	if err != nil {
		return err
	}

	taskFound := false
	for i := 0; i < len(tasks); i++ {
		if tasks[i].ID == taskId {
			taskFound = true
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now().Format(time.RFC3339)
			break
		}

	}

	if !taskFound {
		errorMsg := fmt.Sprintf("Task with the given ID %d not found", taskId)
		return errors.New(errorMsg)
	}

	err = saveTasks(tasks)
	if err != nil {
		return err
	}

	return nil
}

func MarkTask(taskId int, status model.Status) error {
	tasks, err := LoadTasksFromFile()
	if err != nil {
		return err
	}

	taskFound := false
	for i := 0; i < len(tasks); i++ {
		if tasks[i].ID == taskId {
			taskFound = true
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now().Format(time.RFC3339)
			break
		}

	}

	if !taskFound {
		errorMsg := fmt.Sprintf("Task with the given ID %d not found", taskId)
		return errors.New(errorMsg)
	}

	err = saveTasks(tasks)
	if err != nil {
		return err
	}

	return nil
}

func DeleteTask(taskId int) error {
	tasks, err := LoadTasksFromFile()
	if err != nil {
		return err
	}

	taskFoundId := -1
	for i := 0; i < len(tasks); i++ {
		if tasks[i].ID == taskId {
			taskFoundId = i
			break
		}
	}

	if taskFoundId == -1 {
		errorMsg := fmt.Sprintf("Task with the given ID %d not found", taskId)
		return errors.New(errorMsg)
	}

	tasks = append(tasks[:taskFoundId], tasks[taskFoundId+1:]...)

	err = saveTasks(tasks)
	if err != nil {
		return err
	}

	return nil
}
