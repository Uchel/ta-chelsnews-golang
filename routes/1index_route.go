package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func IndexRoute(app *fiber.App, db *gorm.DB) {
	NewsCategoryRoute(app, db)
	NewsRoute(app, db)
	NewsSubRoute(app, db)
}
