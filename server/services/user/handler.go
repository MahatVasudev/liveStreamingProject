package user

import "github.com/gofiber/fiber/v2"

func (h *Handler) GetUserDetails(c *fiber.Ctx) error {

	return c.JSON(map[string]string{"data": "your details..."})
}
