package main

import (
	"fmt"
	"net/http"

	"github.com/Anikcsejnu/ReservationWebApp/pkg/handlers"
)

const portNumber = ":8080"

// main in the main application function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
