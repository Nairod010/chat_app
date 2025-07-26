package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nairod010/chat_app/internal/database"
	"github.com/nairod010/chat_app/internal/models"
)

type APIServer struct {
	listenAddr string
	service    database.Service
}

func NewAPIServer(listenAddr string, service database.Service) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		service:    service,
	}
}

func (s *APIServer) Server() {
	app := fiber.New()
	app.Get("/test", getTest(s.service))
	app.Post("/test", postTest(s.service))
	app.Listen(s.listenAddr)
}

func getTest(service database.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		test, err := service.GetTest()
		if err != nil {
			return c.Status(404).JSON(fiber.Map{
				"erorr": "aia e frate se mai intampla",
			})
		}
		return c.JSON(test)
	}
}

func postTest(service database.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var test models.Test
		if err := c.BodyParser(&test); err != nil {
			return err
		}

		if test.Check == "" {
			return c.Status(400).JSON(fiber.Map{"error": "N-am trimis sefu"})
		}

		if err := service.InsertTest(&test); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "N-am bagat bine"})
		}

		return c.Status(201).JSON(test)
	}
}
