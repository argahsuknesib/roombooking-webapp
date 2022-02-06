package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	fmt.Println(fmt.Println("Starting application on ", portNumber))
	http.ListenAndServe(portNumber, nil)
}
