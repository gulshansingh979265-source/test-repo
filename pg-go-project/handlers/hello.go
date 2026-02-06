package handlers

import (
	"fmt"
	"net/http"
)

// HelloHandler responds with a simple welcome message.
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Go server!")
}

