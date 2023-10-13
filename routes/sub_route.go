package routes

import (
	"chelsnews/handlers"
	"chelsnews/repository"
	"chelsnews/usecase"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewsSubRoute(app *fiber.App, db *gorm.DB) {
	// api.Get("/get-all",	handlers.NewsGetAll)
	// api.Post("/get-all",	handlers.NewsGetAll)
	// api.Get("/get-all",	handlers.NewsGetAll)

	newsSubRepo := repository.NewNewsSubRepository(db)
	newsCtgUseCase := usecase.NewNewsSubUseCase(newsSubRepo)
	newsSubHandler := handlers.NewNewsSubHandler(newsCtgUseCase)

	api := app.Group("/chelsnews/news-sub")
	api.Post("/", newsSubHandler.CreateNews)
	api.Get("/", newsSubHandler.GetNewsCategoryList)
	api.Put("/:name", newsSubHandler.GetNewsCategoryByName)
	api.Delete("/:name", newsSubHandler.DeleteNewsCategory)

}
