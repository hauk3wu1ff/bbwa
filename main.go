package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the handler for the home page
func Home(w http.ResponseWriter, r *http.Request) 
{
	fmt.Fprintf(w, "This is the home page")
}

// About is the handler for the about page
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, "This is the about page and 2 + 2 is %d", sum)
}

// addValues adds two ints x and y, and returns the sum 
func addValues(x, y int) int {
	return x + y
}

// main is the main function
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}