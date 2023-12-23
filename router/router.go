package router

import (
	"net/http"
	"url-shortener/model"

	"github.com/gofiber/fiber/v2"
)

func UrlShortenerRouter(router fiber.Router, us *model.UrlShortener) {
	router.Get("/short/:url", func(c *fiber.Ctx) error {
		rurl := c.Params("url")
		if rurl == "" {
			c.Status(http.StatusBadRequest)
			return c.SendString("URL not found")
		}

		mainURL := us.GetMainURL(rurl)

		if mainURL == "" {
			c.Status(http.StatusBadRequest)
			return c.SendString("URL not found")
		}

		return c.Redirect(mainURL, 302)
	})

	router.Post("/", func(c *fiber.Ctx) error {
		url := c.Query("url")
		if url == "" {
			c.Status(http.StatusBadRequest)
			return c.Send([]byte{})
		}

		data, err := us.ShortURL(url)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.Send([]byte{})
		}
		return c.SendString(data)
	})
}
