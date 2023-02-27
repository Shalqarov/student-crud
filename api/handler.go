package api

import (
	"net/http"
	"net/mail"
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
	if _, err := mail.ParseAddress(student.Email); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"errors": err.Error(),
			})
	}
	err = h.repo.Create(student)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"errors": err.Error(),
			})
	}
	return c.Status(200).JSON(&fiber.Map{
		"student": student,
	})
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
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"student": student,
	})
}

func (h *Handler) readCollection(c *fiber.Ctx) error {
	students, _ := h.repo.FindAll()
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"students": students,
	})
}

func (h *Handler) update(c *fiber.Ctx) error {
	updated := new(repository.Student)
	err := c.BodyParser(updated)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id <= 0 {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	updated.ID = uint64(id)
	student, err := h.repo.Update(updated)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(&fiber.Map{
			"error": "student not found",
		})
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"student": student,
	})
}

func (h *Handler) delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id <= 0 {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	err = h.repo.Delete(uint64(id))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(&fiber.Map{
			"error": "student not found",
		})
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "student deleted successfully",
	})
}
