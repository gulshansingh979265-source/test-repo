// package main


// func main() {
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

	package main

import "fmt"

// Step 1: Define a structure
type Student struct {
    name  string
    age   int
    marks float64
}

func main() {
    // Step 2: Create a variable of struct type
    var s1 Student

    // Step 3: Assign values to struct fields
    s1.name = "Gulshan"
    s1.age = 20
    s1.marks = 85.5

    // Step 4: Print struct values
    fmt.Println("Name:", s1.name)
    fmt.Println("Age:", s1.age)
    fmt.Println("Marks:", s1.marks)
}


	// creating a structure
	// package main
	// import "fmt"
	// type Employee struct {
	// 	name   string
	// 	salary int
	// 	age    int
	// 	empid  string
	// }

	// // creating structure variable
	//  func main(){
	// 	var e Employee
	// 	e.name  = "ANJALI"
	// 	e.salary = 52000
	// 	e.age = 20
	// 	e.empid = "BTECH05"

	// 	fmt.Println(e)

	//  }

	//TODO Implement loop to print struct

	// define structure
	// type Student struct {
	//     name  string
	//     age   int
	//     marks int
	// }

	//     // slice of structures
	//     students := []Student{
	//         {"Gulshan", 20, 85},
	//         {"Amit", 21, 78},
	//         {"Ravi", 22, 90},
	//     }

	//     // loop to print each structure
	//     for i := 0; i < len(students); i++ {
	//         fmt.Println("Student", i+1)
	//         fmt.Println("Name :", students[i].name)
	//         fmt.Println("Age  :", students[i].age)
	//         fmt.Println("Marks:", students[i].marks)
	//         fmt.Println("-----------")
	//     }
	// }

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

