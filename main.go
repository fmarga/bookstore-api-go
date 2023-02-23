package main

import (
	"fmt"

	"github.com/fmarga/bookstore-api-go/routes"
)

func main() {

	fmt.Println("iniciando projeto")
	routes.HandleRequest()
}
