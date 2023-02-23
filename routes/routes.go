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
	log.Fatal(http.ListenAndServe(":8000", r))
}
