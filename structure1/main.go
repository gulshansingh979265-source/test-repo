package main

import (
	"fmt"
	"net/http"
)

// Step 1: Define a structure
type Student struct {
	RollNo int
	Name   string
	Marks  float64
}

// Step 2: Handler function
func studentHandler(w http.ResponseWriter, r *http.Request) {

	// Step 3: Create struct variable
	s := Student{
		RollNo: 101,
		Name:   "Gulshan",
		Marks:  88.5,
	}

	// Step 4: Display struct data on browser
	fmt.Fprintln(w, "Student Details")
	fmt.Fprintln(w, "Roll No:", s.RollNo)
	fmt.Fprintln(w, "Name:", s.Name)
	fmt.Fprintln(w, "Marks:", s.Marks)
}

func main() {

	// Step 5: URL mapping
	http.HandleFunc("/student", studentHandler)

	// Step 6: Start server
	http.ListenAndServe(":8080", nil)
}
