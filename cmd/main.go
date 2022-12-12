package main

import (
	"fmt"
	"net/http"

	"github.com/shubham/bookstore/pkg/models"
	"github.com/shubham/bookstore/pkg/routes"
)

func main() {

	models.ConnectDatabase()
	routes.SetupRoutes()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	http.ListenAndServe(":3000", nil)
	fmt.Println("Server started on port 3000")

}
