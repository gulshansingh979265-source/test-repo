package main

import (
	"example/employee"
)

func main() {

	var emp []employee.Employee
	emp = employee.CreateEmployeeData()
	employee.Display(emp)

}
