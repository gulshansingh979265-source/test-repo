//package student
//
//// Student represents a student record in the system/Postgres.
//type Student struct {
//	ID   int    `json:"id"`
//	Name string `json:"name"`
//	Age  int    `json:"age"`
//}

package student

import _ "fmt"

// Student represents a student record in the system/Postgres.
type Student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Student has 3 generic type parameters:
// IDT   - type of the student ID (string, int, etc.)
// GradT - type of grade/marks (int, float64, etc.)
// MetaT - any extra metadata (struct/map/etc.)
//type Student[IDT any, GradT any, MetaT any] struct {
//	ID     IDT
//	Name   string
//	Grade  GradT
//	Active bool
//	Meta   MetaT
//}

// // Create is a simple constructor-like function.
// func Create[IDT any, GradT any, MetaT any](
// 	id IDT,
// 	name string,
// 	grade GradT,
// 	active bool,
// 	meta MetaT,
// ) Student[IDT, GradT, MetaT] {
// 	return Student[IDT, GradT, MetaT]{
// 		ID:     id,
// 		Name:   name,
// 		Grade:  grade,
// 		Active: active,
// 		Meta:   meta,
// 	}
// }

// // Update updates a few fields on the student.
// func Update[IDT any, GradT any, MetaT any](
// 	s *Student[IDT, GradT, MetaT],
// 	name string,
// 	grade GradT,
// 	active bool,
// 	meta MetaT,
// ) {
// 	s.Name = name
// 	s.Grade = grade
// 	s.Active = active
// 	s.Meta = meta
// }

// // Display prints the student in a simple format.
// func Display[IDT any, GradT any, MetaT any](s Student[IDT, GradT, MetaT]) {
// 	fmt.Printf("ID: %v\n", s.ID)
// 	fmt.Printf("Name: %s\n", s.Name)
// 	fmt.Printf("Grade: %v\n", s.Grade)
// 	fmt.Printf("Active: %v\n", s.Active)
// 	fmt.Printf("Meta: %#v\n", s.Meta)
// 	fmt.Println("----")
// }
