package routes

import (
	"net/http"

	"github.com/shubham/bookstore/pkg/controllers"
)

func SetupRoutes() {
	http.HandleFunc("/books/create", controllers.CreateBook)
	http.HandleFunc("/books", controllers.GetBooks)
	http.HandleFunc("/book", controllers.GetBookById)
	http.HandleFunc("/books/update", controllers.UpdateBook)
	http.HandleFunc("/books/delete", controllers.DeleteBook)
}
