package utils

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func FileCreate(ctx *fiber.Ctx) error {
	//Handle File ==
	file, errFile := ctx.FormFile("image")

	if errFile != nil {
		log.Println(errFile)
	}
	var filename *string
	if file != nil {
		filename = &file.Filename
		errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/images/%s", *filename))
		if errSaveFile != nil {
			log.Println("Failed to save file", errSaveFile)
			return ctx.Status(500).JSON(fiber.Map{
				"message": "failed",
				"error":   "internal server error",
			})
		}

	} else {
		log.Println("no file uploaded")
	}

	if filename != nil {
		ctx.Locals("filename", *filename)
	} else {

		ctx.Locals("filename", nil)
	}

	return ctx.Next()
}
