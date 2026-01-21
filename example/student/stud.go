package student

import "fmt"

type Student struct {
	Name   string
	Salary int
	Age    int
	StudID string
}

func CreateStudentData() []Student {

	students := []Student{
		{Name: "Rahul", Salary: 85, Age: 20, StudID: "ST-001"},
		{Name: "Aman", Salary: 90, Age: 21, StudID: "ST-002"},
		{Name: "Neha", Salary: 88, Age: 19, StudID: "ST-003"},
	}

	return students
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
	// create a multiple  Student (for a slice)
	stud1 := Student{
		Name:   "Rohan",
		Salary: 75000,
		Age:    23,
		StudID: "ST-001",}

		stud2 := Student{
		Name:   "Kishan",
		Salary: 95000,
		Age:    18,
		StudID: "ST-002",
	}

	stud3 := Student{
		Name:   "Balgovind",
		Salary: 950000,
		Age:    20,
		StudID: "ST-003",
	}
	
	


	// append the student
	students = append(students, stud1)
	students = append(students, stud2)
	students = append(students, stud3)

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

func UpdateStudent(students []Student, name string, newSalary int, newAge int) []Student {

	for i, s := range students {
		if s.Name == name {
			students[i].Salary = newSalary
			students[i].Age = newAge
		}
	}

	return students
}

