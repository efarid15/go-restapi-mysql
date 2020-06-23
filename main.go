package main

import (
	"gorestapi/employee"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/employees", employee.Employees)

	println("Listen Port 8008")
	err := http.ListenAndServe(":8008", nil)
	if err != nil {
		log.Fatal(err)

	}
}
