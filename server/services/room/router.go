package room

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router_name string, router *fiber.App) {
	router.Route(router_name, func(router fiber.Router) {
		router.Get("/", h.Welcome)
		router.Post("/create", h.CreateRoom)
		router.Get("/:uuid/websocket", websocket.New(h.RoomWebsocket, websocket.Config{}))
	})
}
