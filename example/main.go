package main

import "fmt"

func main() {
	//TODO -  create vars

	//TODO -  create Structure

	//TODO Implement loop to print struct

	//TO
	// ============ BASIC SWITCH ============
	fmt.Println("--- Basic Switch ---")
	day := "Friday"
	switch day {
	case "Monday":
		fmt.Println("Start of week 📅")
	case "Friday":
		fmt.Println("Almost weekend 🎉")
	case "Saturday", "Sunday":
		fmt.Println("Weekend! 🏖️")
	default:
		fmt.Println("Midweek day")
	}
}
