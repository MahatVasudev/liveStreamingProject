package user

import (
	"github.com/MahatVasudev/liveStreamingProject/server/typestore"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	store typestore.UserStore
}

func NewHandler(store typestore.UserStore) *Handler {
	return &Handler{
		store,
	}
}
func (h *Handler) RegisterRoutes(route_name string, router *fiber.App) {
	router.Route(route_name, func(router fiber.Router) {
		router.Get("/details", h.GetUserDetails)
	})
}
