package main

import (
	"github.com/gofiber/fiber/v2"

	"url-shortener/model"
	"url-shortener/router"
)

func main() {
	app := fiber.New()
	us := model.NewURLShortener()

	router.UrlShortenerRouter(app.Group("/"), us)

	app.Listen(":3000")
}
