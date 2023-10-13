package main

import (
	"chelsnews/config"
	"chelsnews/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	db := config.DatabaseInit()
	config.RunMigration(db)

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000", // Ganti dengan alamat asal frontend Anda
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	routes.IndexRoute(app, db)

	app.Listen(":8080")
}
