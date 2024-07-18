package main

import (
	"fmt"
	"net/http"

	"github.com/ankush-web-eng/mongo/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	fmt.Println("Server started at localhost:8080")
	http.ListenAndServe("localhost:8080", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		fmt.Printf("Failed to connect to MongoDB: %v\n", err)
		panic(err)
	}
	fmt.Println("Connected to MongoDB")
	return s
}
