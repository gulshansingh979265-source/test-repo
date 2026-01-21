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

    var stud1 []student.Student
    stud = student.CreateStudentData()
    student.Display(stud1)




	println("\n\nbefore add\n\n")
	student.Display(stud)
	var newlist []student.Student
	newlist = student.AddStudent(stud)
	println("\n\nafter add\n\n")
	student.Display(newlist)
	newlist = student.DeleteStudent(newlist, "Rahul")
	println("\n\nafter delete\n\n")
	student.Display(newlist)

	newlist = student.UpdateStudent(newlist, "Neha",1200000,20)
	println("\n\nafter update\n\n")
	student.Display(newlist)

    

}
