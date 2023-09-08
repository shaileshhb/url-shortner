package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/shaileshhb/url-shortener/routes"
)

func registerRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		return
	}

	app := fiber.New()
	app.Use(logger.New())
	registerRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
