package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/hello" {
			http.Error(w, "404 Not Found", http.StatusNotFound)
			return
		}
		if r.Method != "GET" {
			http.Error(w, "505 Method is not Supported", http.StatusMethodNotAllowed)
			return
		}
		w.Write([]byte("Hello Stranger!"))
	})

	http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/form" {
			http.Error(w, "505 Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		if err := r.ParseForm(); err != nil {
			http.Error(w, "500 INternal Server Error", http.StatusInternalServerError)
			return
		}
		fmt.Println("Post request successful!!")
		name := r.FormValue("name")
		email := r.FormValue("email")
		fmt.Println("Name: ", name)
		fmt.Println("Email: ", email)
		w.Write([]byte("Name is: " + name + " Email is: " + email))
	})

	println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
