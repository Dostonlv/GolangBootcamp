package repo

type Students interface {
	GetStudents() []Student
	//GetStudentByID(id int) (*Student, error)
	AddStudent(student Student) error
	//RemoveStudentByID(id int) error
	UpdateStudent(id int)
}
type Subject []struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Marks int    `json:"marks"`
}

type Student struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Grade    int     `json:"grade"`
	Section  string  `json:"section"`
	Course   string  `json:"course"`
	Subjects Subject `json:"subjects"`
}

type StudentManager struct {
	Students []Student
}
