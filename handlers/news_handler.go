package handlers

import (
	"chelsnews/models/entity"
	"chelsnews/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type NewsHandler struct {
	NewsUseCase usecase.NewsUseCase
}

func NewNewsHandler(newsUseCase usecase.NewsUseCase) *NewsHandler {
	return &NewsHandler{NewsUseCase: newsUseCase}
}

func (h *NewsHandler) CreateNews(c *fiber.Ctx) error {
	news := new(entity.News)
	if err := c.BodyParser(news); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	filename := c.Locals("filename")

	if filename == nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "image required",
		})
	}
	news.Image = filename.(string)

	if err := h.NewsUseCase.CreateNews(news); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(news)
}

func (h *NewsHandler) GetNewsList(c *fiber.Ctx) error {
	newsList, err := h.NewsUseCase.GetNewsList()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(newsList)
}

func (h *NewsHandler) GetNewsByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	news, err := h.NewsUseCase.GetNewsByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "News not found"})
	}

	return c.JSON(news)
}

func (h *NewsHandler) UpdateNews(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	news := new(entity.News)
	if err := c.BodyParser(news); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	filename := c.Locals("filename")

	if filename == nil {
		news.Image = ""
	}

	news.ID = uint(id)
	if err := h.NewsUseCase.UpdateNews(news); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(news)
}

func (h *NewsHandler) DeleteNews(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := h.NewsUseCase.DeleteNews(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	return c.Status(200).JSON(fiber.Map{"message": "news deleted"})
}
