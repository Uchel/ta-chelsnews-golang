package handlers

import (
	"chelsnews/models/entity"
	"chelsnews/usecase"

	"github.com/gofiber/fiber/v2"
)

type NewsCategoryHandler struct {
	NewsCategoryUseCase usecase.NewsCategoryUseCase
}

func NewNewsCategoryHandler(newsCategoryUseCase usecase.NewsCategoryUseCase) *NewsCategoryHandler {
	return &NewsCategoryHandler{NewsCategoryUseCase: newsCategoryUseCase}
}

func (h *NewsCategoryHandler) CreateNews(c *fiber.Ctx) error {
	newsCtg := new(entity.NewsCategory)
	if err := c.BodyParser(newsCtg); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.NewsCategoryUseCase.CreateNewsCategory(newsCtg); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(newsCtg)
}

func (h *NewsCategoryHandler) GetNewsCategoryList(c *fiber.Ctx) error {
	newsCategoryList, err := h.NewsCategoryUseCase.GetNewsCategoryList()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(newsCategoryList)
}

func (h *NewsCategoryHandler) GetNewsCategoryByName(c *fiber.Ctx) error {
	name := c.Params("name")

	newsCtg, err := h.NewsCategoryUseCase.GetNewsCategoryByName(name)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "News not found"})
	}

	return c.JSON(newsCtg)
}
func (h *NewsCategoryHandler) UpdateNewsCategoryByName(c *fiber.Ctx) error {
	name := c.Params("name")
	newsCtg := new(entity.NewsCategory)
	if err := c.BodyParser(newsCtg); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.NewsCategoryUseCase.UpdateNewsCategoryByName(newsCtg, name)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "News not found"})
	}

	return c.JSON(newsCtg)
}
func (h *NewsCategoryHandler) DeleteNewsCategory(c *fiber.Ctx) error {
	name := c.Params("name")
	if err := h.NewsCategoryUseCase.DeleteNewsCategory(name); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{"message": "news deleted"})
}
