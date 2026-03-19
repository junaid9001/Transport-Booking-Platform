package handlers

import "github.com/gofiber/fiber/v3"

// GetLiveStatus handles GET /api/train/:id/live-status
func GetLiveStatus() fiber.Handler {
	return func(c fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"message": "live status - coming in phase 4"})
	}
}
