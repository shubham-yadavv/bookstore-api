package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/shubham/bookstore/pkg/database"
	"github.com/shubham/bookstore/pkg/models"
)

func GetBooks(c *fiber.Ctx) error {
	db := database.DB
	var books []models.Book

	db.Find(&books)

	if len(books) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "No books found",
		})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Books Found", "data": books})
}

func GetBook(c *fiber.Ctx) error {
	db := database.DB
	var book models.Book

	id := c.Params("id")

	db.Find(&book, "id = ?", id)

	if book.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No Book present", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Book Found", "data": book})
}

func CreateBooks(c *fiber.Ctx) error {
	db := database.DB
	book := new(models.Book)

	err := c.BodyParser(book)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	book.ID = uuid.New()

	err = db.Create(&book).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create book", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created Book", "data": book})

}

func UpdateBook(c *fiber.Ctx) error {
	type updateBook struct {
		Title  string `json:"title"`
		Author string `json:"author"`
	}
	db := database.DB
	var book models.Book

	id := c.Params("id")

	db.Find(&book, "id = ?", id)

	if book.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
	}

	var updateBookData updateBook
	err := c.BodyParser(&updateBookData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	book.Title = updateBookData.Title
	book.Author = updateBookData.Author

	db.Save(&book)

	return c.JSON(fiber.Map{"status": "success", "message": "Books Found", "data": book})

}

func DeleteBook(c *fiber.Ctx) error {
	db := database.DB
	var book models.Book

	id := c.Params("id")

	db.Find(&book, "id = ?", id)

	if book.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
	}

	err := db.Delete(&book, "id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete book", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Deleted book"})

}

func DeleteAllBooks(c *fiber.Ctx) error {
	db := database.DB
	var books []models.Book

	db.Find(&books)
	db.Delete(&books)

	return c.JSON(fiber.Map{
		"message": "All books deleted",
	})

}
