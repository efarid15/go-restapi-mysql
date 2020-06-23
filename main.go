package main

import (
	"gorestapi/handlers"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/employees", handlers.Employee)

	println("Listen Port 8008")
	err := http.ListenAndServe(":8008", nil)
	if err != nil {
		log.Fatal(err)

	}
}
