package main

import (
	"fmt"

	"github.com/fmarga/bookstore-api-go/db"
	"github.com/fmarga/bookstore-api-go/routes"
)

func main() {

	db.Conn()
	fmt.Println("iniciando projeto")
	routes.HandleRequest()
}
