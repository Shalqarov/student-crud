package api

import (
	"github.com/Shalqarov/student-crud/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

// Setup receives a connection to the database as input, and passes it to the repository.
// This is also where routers and middleware are configured.
func Setup(db *gorm.DB) *fiber.App {
	h := &Handler{
		repo: repository.NewRepo(db),
	}
	app := *fiber.New()
	app.Use(logger.New())
	app.Post("/api/v1/students", h.create)
	app.Get("/api/v1/students/:id", h.read)
	app.Get("/api/v1/students/", h.readCollection)
	app.Put("/api/v1/students/:id", h.update)
	app.Delete("/api/v1/students/:id", h.delete)
	return &app
}
