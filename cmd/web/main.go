package main

import (
	"fmt"
	"github.com/rozdolsky33/Booking-App/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Println(fmt.Sprintf("Starting application on port: %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
