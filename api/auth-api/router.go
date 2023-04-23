package authAPI

import (
	"github.com/gofiber/fiber/v2"
)

const (
	HandlerPath = "/auth"
)

func RegisterRoutes(router fiber.Router) {
	router.Post("/login", LoginHandler)
}
