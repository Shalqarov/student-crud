package main

import (
	"github.com/Shalqarov/student-crud/api"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := api.Setup()
	app.Use(logger.New())
	app.Listen(":8080")
}
