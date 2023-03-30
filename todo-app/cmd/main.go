package main

import (
	"encoding/json"
	"fmt"
	"os"
	"todo/internal/db"
	"todo/internal/repo"
)

func main() {

	data, err := os.ReadFile("../internal/data/todo.json")
	if err != nil {
		panic(err)
	}
	var task []repo.Task
	err = json.Unmarshal(data, &task)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	Task := &db.TaskManager{Tasks: task}
	var taskid int
START:
	fmt.Println("----------Todo-App-----------")
	fmt.Println("1-Todo list")
	fmt.Println("2-Add task to Todo")
	fmt.Println("3-Delete task to Todo")
	fmt.Println("4- Marked  task's Complete to Todo")
	fmt.Print("Enter task id: ")
	fmt.Scan(&taskid)
	//fmt.Println(Task)
	if taskid == 1 {
		for _, task := range Task.GetAllTasks() {
			fmt.Printf(`----------------------------------
Task ID: %v
Task Title: %v
Compelete task: %v
Description task :%v
Due Date Task: %v
Priority Task: %v
Start Date Task:%v
`, task.ID, task.Title, task.Completed, task.Description, task.DueDate, task.Priority, task.StartDate)
		}
	} else if taskid == 2 {
		Task.AddTask(repo.Task{
			ID:          len(Task.GetAllTasks()) + 1,
			Title:       "Learn Http and Https",
			Completed:   false,
			Description: "how to fixed 500 http code 😁",
			DueDate:     "2023-05-30",
			Priority:    "easy",
			StartDate:   "2023-05-27",
		})
		Task.WriteTasks()

		goto START
	} else if taskid == 3 {
		var id int
		fmt.Print("Enter the task id you want to delete: ")
		fmt.Scan(&id)
		Task.RemoveTask(id)
		Task.WriteTasks()
		goto START
	} else if taskid == 4 {
		var id int
		fmt.Print("Enter the task id you want to mark to complete: ")
		fmt.Scan(&id)
		Task.MarkTaskCompleted(id)
		goto START
	}
}
