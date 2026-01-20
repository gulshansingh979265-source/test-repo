package student

import "fmt"

type Student struct {
	Name   string
	Salary int
	Age    int
	StudID string
}

func Display(students []Student) {
	for i := 0; i < len(students); i++ {
		fmt.Println("Student", i+1)
		fmt.Println("Name  :", students[i].Name)
		fmt.Println("Salary:", students[i].Salary)
		fmt.Println("Age   :", students[i].Age)
		fmt.Println("StudID:", students[i].StudID)
		fmt.Println("-------------------")
	}
}

func AddStudent(students []Student) []Student {
	// create a single Student (not a slice)
	stud := Student{
		Name:   "Rohan",
		Salary: 75000,
		Age:    23,
		StudID: "ST-001",
	}

	// append the student
	students = append(students, stud)

	return students
}

func DeleteStudent(students []Student, name string) []Student {
	for i, s := range students {
		if s.Name == name {
			// remove the element at index i
			students = append(students[:i], students[i+1:]...)
			return students
		}
	}
	return students
}
