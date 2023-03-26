package main

import (
	"encoding/json"
	"fmt"
	"os"
	"structure/storage/jsonning"
	"structure/storage/repo"
)

func main() {
	data, err := os.ReadFile("../storage/jsonning/student.json")
	if err != nil {
		panic(err)
	}
	var student []repo.Student
	err = json.Unmarshal(data, &student)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	Std := &jsonning.StudentManager{Students: student}

START:
	var a int
	fmt.Println("----------Welcome to University-----------")
	fmt.Println("1-Studentlar listi")
	fmt.Println("2-Student qo'shish ")
	fmt.Println("3-Student Update ")
	fmt.Println("4-Studentni ID orqali olish ")
	fmt.Println("5-Studentni ID orqali 0'chirish ")
	fmt.Print("Sonni kiriting: ")
	fmt.Scan(&a)
	if a == 1 {
		for _, std := range Std.Students {
			fmt.Printf(`--------------------------
Student ID: %v,
Student ismi: %v
Student grade: %v
`, std.ID, std.Name, std.Grade)
		}

	} else if a == 2 {
		Std.AddStudent(repo.Student{
			ID:     len(Std.Students),
			Name:   "Jasur",
			Grade:  70,
			Course: "",
			Subjects: repo.Subject{
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
	} else if a == 4 {
		var id int
		fmt.Print("Id ni kiriting: ")
		fmt.Scan(&id)
		student, err := Std.GetStudentByID(id)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf(`--------------------------
Student ID: %v,
Student ismi: %v
Student grade: %v
Student Section: %v


`, student.ID, student.Name, student.Grade, student.Section)
		}
		goto START
	} else if a == 5 {
		var id int
		fmt.Print("Id ni kiriting: ")
		fmt.Scan(&id)
		fmt.Println(Std.RemoveStudentByID(id))
		goto START
	}
}
