package main

import (
	"fmt"
	"net/http"
)

// Home page
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Home Page")
}

// About page
func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is About Page")
}

// Contact page
func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Contact us at: support@example.com")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)

	http.ListenAndServe(":8080", nil)
}
