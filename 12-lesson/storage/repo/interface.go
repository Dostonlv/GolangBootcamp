package repo

type Students interface {
	GetStudents() []Student
	GetStudentByID(id int) (*Student, error)
	AddStudent(student Student) error
	RemoveStudentByID(id int) error
	UpdateStudent(id int)
}
type Subject []struct {
	ID    int    `jsonning:"id"`
	Name  string `jsonning:"name"`
	Marks int    `jsonning:"marks"`
}

type Student struct {
	ID       int     `jsonning:"id"`
	Name     string  `jsonning:"name"`
	Grade    int     `jsonning:"grade"`
	Section  string  `jsonning:"section"`
	Course   string  `jsonning:"course"`
	Subjects Subject `jsonning:"subjects"`
}

type StudentManager struct {
	Students []Student
}
