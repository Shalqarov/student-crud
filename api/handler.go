package api

import (
	"net/http"
	"strconv"

	"github.com/Shalqarov/student-crud/repository"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	repo repository.Repo
}

func (h *Handler) create(c *fiber.Ctx) error {
	student := new(repository.Student)
	err := c.BodyParser(&student)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(
			fiber.Map{
				"errors": err.Error() + "; content-type must be 'application/json'",
			})
	}
	err = h.repo.Create(student)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"errors": err.Error(),
			})
	}
	c.Status(200).JSON(&fiber.Map{
		"student": student,
	})

	return nil
}

func (h *Handler) read(c *fiber.Ctx) error {
	strId := c.Params("id")
	id, err := strconv.Atoi(strId)
	if err != nil || id <= 0 {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": "StatusBadRequest",
		})
	}
	student, err := h.repo.FindOne(uint64(id))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(&fiber.Map{
			"error": "student not found",
		})

	}
	c.Status(http.StatusOK).JSON(&fiber.Map{
		"student": student,
	})
	return nil
}

func (h *Handler) readCollection(c *fiber.Ctx) error {
	students, _ := h.repo.FindAll()
	c.Status(http.StatusOK).JSON(&fiber.Map{
		"students": students,
	})
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
