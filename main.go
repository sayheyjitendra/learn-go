package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello World In Go")
}

func main(){


	http.HandleFunc("/", Home)

	fmt.Printf("Running on Port %s", portNumber)


	_= http.ListenAndServe(portNumber, nil)

}