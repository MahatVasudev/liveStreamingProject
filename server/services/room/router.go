package room

import "github.com/gofiber/fiber/v2"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func RegisterRoutes(router_name string, router *fiber.App) {
	router.Route(router_name, func(router fiber.Router) {
		router.Get("/", nil)

		router.Post("/create", nil)
	})
}
