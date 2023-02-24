package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fmarga/bookstore-api-go/db"
	"github.com/fmarga/bookstore-api-go/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book

	db.DB.Find(&books)
	json.NewEncoder(w).Encode(books)
}

func Book(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var book []models.Book

	db.DB.First(&book, id)
	json.NewEncoder(w).Encode(book)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)
	db.DB.Create(&book)

	json.NewEncoder(w).Encode(&book)
}

func EditBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var book models.Book

	db.DB.First(&book, id)
	json.NewDecoder(r.Body).Decode(&book)
	db.DB.Save(&book)

	json.NewEncoder(w).Encode(&book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var book models.Book

	db.DB.Delete(&book, id)
	json.NewEncoder(w).Encode(book)
}
