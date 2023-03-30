package repo

type ToDoList interface {
	GetAllTasks() []Task
	AddTask(task Task) error
	RemoveTask(index int) error
	MarkTaskCompleted(index int) error
	WriteTasks() error
}

type Task struct {
	ID          int    `data:"id"`
	Title       string `data:"title"`
	Completed   bool   `data:"completed"`
	Description string `data:"description"`
	DueDate     string `data:"dueDate"`
	Priority    string `data:"priority"`
	StartDate   string `data:"startDate"`
}

type TaskManager struct {
	Tasks []Task
}
