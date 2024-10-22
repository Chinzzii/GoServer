package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil{
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprint(w, "POST Request Successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprint(w, "Name = %s\n", name)
	fmt.Fprint(w, "Address = %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w, "Methof is not Supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello from Go Server!")
}

func main() {
	fileServer := http.FileServer(http.Dir("/static"))
	http.Handle("/form", formHandler)
	http.Handle("/hello", helloHandler)

	fmt.Printf("Starting server at 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil{
		log.Fatal(err)
	}
}
