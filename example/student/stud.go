package student

import "fmt"

type Student struct {
	Name   string
	Salary int
	Age    int
	StudID  string
}

    //   var Students = []Student{
	// 	{Name: "Anjali", Salary: 52000, Age: 20, StudID: "BTECH05"},
	// 	{Name: "Rahul", Salary: 60000, Age: 23, StudID: "BTECH06"},
	// 	{Name: "Neha", Salary: 58000, Age: 22, StudID: "BTECH07"},
	// }

func Display(students []Student) {
	for i := 0; i < len(students); i++ {
		fmt.Println("Student", i+1)
		fmt.Println("Name  :", students[i].Name)
		fmt.Println("Salary:", students[i].Salary)
		fmt.Println("Age   :", students[i].Age)
		fmt.Println("StudID :", students[i].StudID)
		fmt.Println("-------------------")
	}
}