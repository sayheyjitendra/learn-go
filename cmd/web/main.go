package main

import (
	"fmt"
	"net/http"
	"github.com/sayheyjitendra/learn-go/package/handlers"
)

const portNumber = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	
	fmt.Printf("Running on Port %s", portNumber)

	_= http.ListenAndServe(portNumber, nil)
	
}