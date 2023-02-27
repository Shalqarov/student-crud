package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Setup() *fiber.App {
	m := &StudentModel{}
	app := *fiber.New()
	app.Use(logger.New())
	app.Post("/api/v1/students", m.create)
	app.Get("/api/v1/students/:id", m.read)
	app.Get("/api/v1/students/", m.readCollection)
	app.Put("/api/v1/students/:id", m.update)
	app.Delete("/api/v1/students/:id", m.delete)
	return &app
}
