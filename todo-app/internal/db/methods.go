package db

import (
	"encoding/json"
	"fmt"
	"os"
	"todo/internal/repo"
)

type TaskManager repo.TaskManager

func (t *TaskManager) GetAllTasks() []repo.Task {
	return t.Tasks
}

func (t *TaskManager) AddTask(task repo.Task) error {

	t.Tasks = append(t.Tasks, task)
	return nil
}

func (t *TaskManager) RemoveTask(id int) error {
	for i, task := range t.Tasks {
		if task.ID == id {
			t.Tasks = append(t.Tasks[:i], t.Tasks[i+1:]...)
			return fmt.Errorf("task removed successfully id:%v", id)
		}
	}
	return fmt.Errorf("task id %v is not found", id)
}

func (t *TaskManager) MarkTaskCompleted(id int) error {
	data, err := os.ReadFile("../internal/data/todo.json")
	if err != nil {
		panic(err)
	}
	var task []repo.Task

	err = json.Unmarshal(data, &task)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
	}
	var tasks repo.Task
	var taska []repo.Task
	for _, v := range task {
		if v.ID == id {
			tasks = v
		}
	}
	tasks.Completed = true
	taska = append(taska, tasks)
	for _, v := range task {
		if v.ID != id {
			taska = append(taska, v)
		}
	}
	jsonData, err := json.Marshal(taska)
	if err != nil {
		return fmt.Errorf("failed to marshal tasks to JSON: %v", err)
	}
	err = os.WriteFile("../internal/data/todo.json", jsonData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write tasks to JSON file: %v", err)
	}

	return fmt.Errorf("task id %v not found", id)
}

func (t *TaskManager) WriteTasks() error {
	jsonData, err := json.Marshal(t.Tasks)
	if err != nil {
		return fmt.Errorf("failed to marshal tasks to JSON: %v", err)
	}

	// Write the JSON data to the file
	err = os.WriteFile("../internal/data/todo.json", jsonData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write tasks to JSON file: %v", err)
	}

	return nil
}
