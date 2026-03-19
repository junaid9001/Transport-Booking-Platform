package handlers

import "github.com/gofiber/fiber/v3"

func SearchTrains() fiber.Handler {
	return func(c fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"message": "search - coming in phase 4"})
	}
}

func GetTrainByID() fiber.Handler {
	return func(c fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"message": "get train - coming in phase 4"})
	}
}
