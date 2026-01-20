package main

import (
	//"example/employee"
	"example/student"
)

func main() {

	 stud := []student.Student{
		{Name: "Anjali", Salary: 52000, Age: 20, StudID: "BTECH05"},
		{Name: "Rahul", Salary: 60000, Age: 23, StudID: "BTECH06"},
		{Name: "Neha", Salary: 58000, Age: 22, StudID: "BTECH07"},
	}

	// var emp []employee.Employee
	// emp = employee.CreateEmployeeData()
	// employee.Display(emp)

	//  var stud []student.Student
	// emp = employee.CreateEmployeeData()
	 student.Display(stud)
  



}
