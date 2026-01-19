package employee

import "fmt"

type Employee struct {
	Name   string
	Salary int
	Age    int
	EmpID  string
}

func CreateEmployeeData() []Employee {
	employees := []Employee{
		{Name: "Anjali", Salary: 52000, Age: 20, EmpID: "BTECH05"},
		{Name: "Rahul", Salary: 60000, Age: 23, EmpID: "BTECH06"},
		{Name: "Neha", Salary: 58000, Age: 22, EmpID: "BTECH07"},
	}
	return employees
}

func Display(employees []Employee) {
	for i := 0; i < len(employees); i++ {
		fmt.Println("Employee", i+1)
		fmt.Println("Name  :", employees[i].Name)
		fmt.Println("Salary:", employees[i].Salary)
		fmt.Println("Age   :", employees[i].Age)
		fmt.Println("EmpID :", employees[i].EmpID)
		fmt.Println("-------------------")
	}
}
