package jsonning

import (
	"structure/storage/repo"
)

var a int

type StudentManager repo.StudentManager

func (s *StudentManager) AddStudent(student repo.Student) error {
	s.Students = append(s.Students, student)
	return nil
}

func (s *StudentManager) GetStudents() []repo.Student {
	return s.Students
}

//
//func (s *repo.StudentManager) UpdateStudent(id int) {
//	for i, std := range s.students {
//		if std.ID == id {
//			fmt.Printf("Old Name: %s\n", std.Name)
//			fmt.Print("Enter new Name: ")
//			fmt.Scan(&std.Name)
//			fmt.Printf("Old Grade: %d\n", std.Grade)
//			fmt.Print("Enter new Grade: ")
//			fmt.Scan(&std.Grade)
//			s.students[i] = std
//			fmt.Println("Student updated successfully!")
//			return
//		}
//	}
//	fmt.Printf("Student with ID %d not found!\n", id)
//}
