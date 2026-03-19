package middleware

import "github.com/gofiber/fiber/v3"

// ExtractUser reads X-User-ID and X-User-Role headers injected by the API Gateway
// after it validates the JWT. The train service trusts these headers.
func ExtractUser() fiber.Handler {
	return func(c fiber.Ctx) error {
		userID := c.Get("X-User-ID")
		if userID == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "unauthorized",
			})
		}
		c.Locals("userID", userID)
		c.Locals("userRole", c.Get("X-User-Role"))
		return c.Next()
	}
}
