package server

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

func ctxWithTimeout(c *fiber.Ctx) (context.Context, context.CancelFunc) {
	return context.WithTimeout(c.Context(), 10*time.Second)
}
