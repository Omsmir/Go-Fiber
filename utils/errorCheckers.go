package utils

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func ErrorCheck(msg string, err error) {
	if err != nil {
		log.Fatalf("❌ %s: %v", msg, err)
	}
}

func StatusDetector(c *fiber.Ctx, err *fiber.Error) error {
	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"error": err.Message,
		})
	}
	return nil
}
