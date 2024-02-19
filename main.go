package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 3)
	_, _ = fmt.Fprint(w, fmt.Sprintf("This is the about page and 2 + 3 is %d", sum))
}

// AddValue add two integer and return sum
func addValues(x, y int) int {
	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValue(100.0, 0.0)

	if err != nil {
		fmt.Fprint(w, "Can not devide by 0")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%f divided bt %f = %f", 100.0, 10.0, f))

}

func divideValue(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("Can not divide by zero")
		return 0, err
	}

	result := x / y
	return result, nil
}

// main in the main application function
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
