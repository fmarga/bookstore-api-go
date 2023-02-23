package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fmarga/bookstore-api-go/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	json.NewEncoder(w).Encode(books)
}
