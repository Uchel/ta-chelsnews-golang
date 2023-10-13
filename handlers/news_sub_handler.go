package handlers

import (
	"chelsnews/models/entity"
	"chelsnews/usecase"

	"github.com/gofiber/fiber/v2"
)

type NewsSubHandler struct {
	NewsSubUseCase usecase.NewsSubUseCase
}

func NewNewsSubHandler(newsSubUseCase usecase.NewsSubUseCase) *NewsSubHandler {
	return &NewsSubHandler{NewsSubUseCase: newsSubUseCase}
}

func (h *NewsSubHandler) CreateNews(c *fiber.Ctx) error {
	newsSub := new(entity.SubCategory)
	if err := c.BodyParser(newsSub); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.NewsSubUseCase.CreateNewsCategory(newsSub); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(newsSub)
}

func (h *NewsSubHandler) GetNewsCategoryList(c *fiber.Ctx) error {
	newsSubList, err := h.NewsSubUseCase.GetNewsCategoryList()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(newsSubList)
}

func (h *NewsSubHandler) GetNewsCategoryByName(c *fiber.Ctx) error {
	name := c.Params("name")

	newsSub, err := h.NewsSubUseCase.GetNewsCategoryByName(name)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "News not found"})
	}

	return c.JSON(newsSub)
}

func (h *NewsSubHandler) DeleteNewsCategory(c *fiber.Ctx) error {
	name := c.Params("name")
	if err := h.NewsSubUseCase.DeleteNewsCategory(name); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{"message": "news deleted"})
}
