package routes

import (
	"log"
	"net/http"

	"github.com/fmarga/bookstore-api-go/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/all", controllers.AllBooks).Methods("Get")
	r.HandleFunc("/api/all/{id}", controllers.Book).Methods("Get")
	r.HandleFunc("/api/all", controllers.AddBook).Methods("Post")
	r.HandleFunc("/api/all/{id}", controllers.EditBook).Methods("Put")
	r.HandleFunc("/api/all/{id}", controllers.DeleteBook).Methods("Delete")
	log.Fatal(http.ListenAndServe(":8000", r))
}
