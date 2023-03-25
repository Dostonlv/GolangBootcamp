package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Students interface {
	GetStudents() []Student
	//GetStudentByID(id int) (*Student, error)
	AddStudent(student Student) error
	//RemoveStudentByID(id int) error
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
	students []Student
}

func (s *StudentManager) AddStudent(student Student) error {
	s.students = append(s.students, student)
	return nil
}

func (s *StudentManager) GetStudents() []Student {
	return s.students
}

func (s *StudentManager) UpdateStudent(id int) {
	for i, std := range s.students {
		if std.ID == id {
			fmt.Printf("Old Name: %s\n", std.Name)
			fmt.Print("Enter new Name: ")
			fmt.Scan(&std.Name)
			fmt.Printf("Old Grade: %d\n", std.Grade)
			fmt.Print("Enter new Grade: ")
			fmt.Scan(&std.Grade)
			s.students[i] = std
			fmt.Println("Student updated successfully!")
			return
		}
	}
	fmt.Printf("Student with ID %d not found!\n", id)
}

func main() {
	data, err := os.ReadFile("student.jsonning")
	if err != nil {
		panic(err)
	}
	var student []Student
	err = json.Unmarshal(data, &student)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	Std := &StudentManager{student}

START:
	var a int
	fmt.Println("----------Welcome to University-----------")
	fmt.Println("1-Studentlar listi")
	fmt.Println("2-Student qo'shish ")
	fmt.Println("3-Student Update ")
	fmt.Print("Sonni kiriting: ")
	fmt.Scan(&a)
	if a == 1 {
		for _, std := range Std.GetStudents() {
			fmt.Printf(`--------------------------
Student ID: %v,
Student ismi: %v
Student grade: %v
`, std.ID, std.Name, std.Grade)
		}

	} else if a == 2 {
		Std.AddStudent(Student{
			ID:     len(Std.GetStudents()),
			Name:   "Jasur",
			Grade:  70,
			Course: "",
			Subjects: Subject{
				{
					ID:    1,
					Name:  "Maths",
					Marks: 70,
				},
				{
					ID:    2,
					Name:  "Science",
					Marks: 70,
				},
				{
					ID:    3,
					Name:  "English",
					Marks: 70,
				},
			},
		})

		goto START
	} else if a == 3 {
		var id int
		fmt.Print("Id ni kiriting: ")
		fmt.Scan(&id)
		Std.UpdateStudent(id)
		goto START
	}
}
