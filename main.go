package main

import (
	"fmt"
	"go-quickstart/to-do-list/operation"
	"net/http"
)

func main() {

	http.HandleFunc("/", operation.Sample)
	http.HandleFunc("/get/alltask", operation.GetAllTask)
	http.HandleFunc("/add/task", operation.Add)
	fmt.Println("here")

	http.ListenAndServe(":8080", nil)

}
