package main

import "fmt"

func main() {
	//TODO -  create vars
//  {
//     var age int = 20
//     var name string = "Gulshan"

//     fmt.Println(age)
//     fmt.Println(name)
// }


//  {
//     city := "Delhi"
//     marks := 85

//     fmt.Println(city)
//     fmt.Println(marks)
// }

	//TODO -  create Structure
	

// creating a structure
// type Student struct {
//     name  string
//     age   int
//     marks int
// }

//  {
//     // creating structure variable
//     var s Student

//     s.name = "Gulshan"
//     s.age = 20
//     s.marks = 85

//     fmt.Println(s)
// }

 // creating a structure 
 type Employee struct{
	name string
	salary int
	age int
	empid string
 }

 // creating structure variable
 {
	var e Employee
	e.name  = "ANJALI"
	e.salary = 52000
	e.age = 20
	e.empid = "BTECH05"

	fmt.Println(e)

 }





	//TODO Implement loop to print struct

	//TO
	// ============ BASIC SWITCH ============
// 	fmt.Println("--- Basic Switch ---")
// 	day := "Monday"
// 	switch day {
// 	case "Monday":
// 		fmt.Println("Start of week 📅")
// 	case "Friday":
// 		fmt.Println("Almost weekend 🎉")
// 	case "Saturday", "Sunday":
// 		fmt.Println("Weekend! 🏖️")
// 	default:
// 		fmt.Println("Midweek day")
// 	}
}
