package api

import (
	"github.com/Shalqarov/student-crud/repository"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	repo repository.Repo
}

func (h *Handler) create(c *fiber.Ctx) error {
	// TODO
	return nil
}

func (h *Handler) read(c *fiber.Ctx) error {
	// TODO
	return nil
}

func (h *Handler) readCollection(c *fiber.Ctx) error {
	// TODO
	return nil
}

func (h *Handler) update(c *fiber.Ctx) error {
	// TODO
	return nil
}

func (h *Handler) delete(c *fiber.Ctx) error {
	// TODO
	return nil
}
