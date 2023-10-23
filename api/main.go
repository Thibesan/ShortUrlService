package main

import (
	"fmt"
	"log"
	"os"

	"github.com/goFiber/fiber/v2"
	"github.com/goFiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
}

func main() {
	err := godotenv.Load()

	if err!=nil {
		fmt.println(err)
	}

	app := fiber.New()
	app.use(logger.New())
	setupRoutes(app)
	log.Fatal(app.listen(os.Getenv("APP_PORT")))
}