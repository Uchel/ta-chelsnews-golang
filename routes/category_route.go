package routes

import (
	"chelsnews/handlers"
	"chelsnews/repository"
	"chelsnews/usecase"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewsCategoryRoute(app *fiber.App, db *gorm.DB) {
	// api.Get("/get-all",	handlers.NewsGetAll)
	// api.Post("/get-all",	handlers.NewsGetAll)
	// api.Get("/get-all",	handlers.NewsGetAll)

	newsCtgRepo := repository.NewNewsCategoryRepository(db)
	newsCtgUseCase := usecase.NewNewsCategoryUseCase(newsCtgRepo)
	newsCtgHandler := handlers.NewNewsCategoryHandler(newsCtgUseCase)

	api := app.Group("/chelsnews/news-ctg")
	api.Post("/", newsCtgHandler.CreateNews)
	api.Get("/", newsCtgHandler.GetNewsCategoryList)
	api.Get("/:name", newsCtgHandler.GetNewsCategoryByName)
	api.Put("/:name", newsCtgHandler.UpdateNewsCategoryByName)
	api.Delete("/:name", newsCtgHandler.DeleteNewsCategory)

}
