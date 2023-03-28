package repo

type ToDoList interface {
	GetAllTasks() []Task
	AddTask(task Task) error
	RemoveTask(index int) error
	MarkTaskCompleted(index int) error
	WriteTasks() error
}

type Task struct {
	ID          int    `db:"id"`
	Title       string `db:"title"`
	Completed   bool   `db:"completed"`
	Description string `db:"description"`
	DueDate     string `db:"dueDate"`
	Priority    string `db:"priority"`
	StartDate   string `db:"startDate"`
}

type TaskManager struct {
	Tasks []Task
}
