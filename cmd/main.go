package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shubham/bookstore/pkg/routes"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:3000", r))
	fmt.Println("Server is running on port 3000")
}
