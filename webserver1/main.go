package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, " my name is gulshan singh")

}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "  jmhkjhljhjv")

}

func handler3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "  ttdkndckcn")

}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/test", handler2)
	http.HandleFunc("/test1", handler3)
	http.ListenAndServe(":8080", nil)

}
