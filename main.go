package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()
	app.Use(recover.New(recover.Config{EnableStackTrace: true}))
	app.Use(logger.New())

	app.Get("/api/list", ListOpenableLab)
	app.Get("/api/open", OpenDoor)

	_ = app.Listen("0.0.0.0:8000")
}
