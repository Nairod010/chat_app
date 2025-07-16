package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nairod010/chat_app/internal/database"
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
