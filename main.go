package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		fmt.Fprintf(res, "Parse form error: %v", err)
		return
	}
	fmt.Fprintf(res, "Post request successful\n")
	name := req.FormValue("name")
	address := req.FormValue("address")
	fmt.Fprintf(res, "Name: %s, Address: %s. \n", name, address)
}

func helloHandler(res http.ResponseWriter, req *http.Request) {

	if req.URL.Path != "/hello" {
		http.Error(res, "404 not found", http.StatusNotFound)
		return
	}

	if req.Method != "GET" {
		http.Error(res, "Mesthod not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(res, "Hello Go!")
}

func main() {
	const PORT int = 8080
	fileserver := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("starting server at port ", PORT)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
