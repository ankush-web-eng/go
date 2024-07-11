package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ankush-web-eng/Bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookstoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Server started at localhost:9010")
	log.Fatal(http.ListenAndServe("localhost:9010", nil))
}
