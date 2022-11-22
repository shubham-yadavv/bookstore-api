package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/shubham/bookstore/pkg/config"
	"github.com/shubham/bookstore/pkg/database"
	"github.com/shubham/bookstore/pkg/routes"
)

func main() {
	app := fiber.New()

	database.ConnectDB()

	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("And the API is UP!")
		return err
	})

	routes.SetupRoutes(app)

	port := config.Config("PORT")

	fmt.Println("Server is running on port " + port)
	app.Listen(":" + port)

}
