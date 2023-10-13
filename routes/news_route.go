package routes

import (
	"chelsnews/config"
	"chelsnews/handlers"
	"chelsnews/repository"
	"chelsnews/usecase"
	"chelsnews/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewsRoute(app *fiber.App, db *gorm.DB) {
	// api.Get("/get-all",	handlers.NewsGetAll)
	// api.Post("/get-all",	handlers.NewsGetAll)
	// api.Get("/get-all",	handlers.NewsGetAll)

	newsRepo := repository.NewNewsRepository(db)
	newsUseCase := usecase.NewNewsUseCase(newsRepo)
	newsHandler := handlers.NewNewsHandler(newsUseCase)

	api := app.Group("/chelsnews/news")
	api.Static("/photo", config.ProjectRoutePath+"/public/images")
	api.Post("/", utils.FileCreate, newsHandler.CreateNews)
	api.Get("/", newsHandler.GetNewsList)
	api.Get("/:id", newsHandler.GetNewsByID)
	api.Put("/:id", utils.FileCreate, newsHandler.UpdateNews)
	api.Delete("/:id", newsHandler.DeleteNews)

}
