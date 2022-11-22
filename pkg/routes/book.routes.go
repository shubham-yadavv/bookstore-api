package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shubham/bookstore/pkg/controllers"
)

func SetupRoutes(router fiber.Router) {
	books := router.Group("/books")

	books.Post("/", controllers.CreateBook)
	books.Get("/", controllers.GetBooks)
	books.Get("/:id", controllers.GetBook)
	books.Put("/:id", controllers.UpdateBook)
	books.Delete("/:id", controllers.DeleteBook)
}
