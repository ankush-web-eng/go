package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ankush-web-eng/Bookstore/pkg/models"
	"github.com/ankush-web-eng/Bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

// var newBook models.Book

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookdId := params["bookId"]
	ID, err := strconv.ParseInt(bookdId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing Id")
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	Createbook := &models.Book{}
	utils.ParseBody(r, Createbook)
	b := Createbook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookdId := params["bookId"]
	ID, err := strconv.ParseInt(bookdId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing Id")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	UpdateBook := &models.Book{}
	utils.ParseBody(r, UpdateBook)
	params := mux.Vars(r)
	bookdId := params["bookId"]
	ID, err := strconv.ParseInt(bookdId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing Id")
	}
	bookDetails, db := models.GetBookById(ID)
	if UpdateBook.Name != "" {
		bookDetails.Name = UpdateBook.Name
	}
	if UpdateBook.Author != "" {
		bookDetails.Author = UpdateBook.Author
	}
	if UpdateBook.Publication != "" {
		bookDetails.Publication = UpdateBook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
