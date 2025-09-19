package server

import (
	"github.com/candra/saas-rs-backend/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func createInvoice(pool *pgxpool.Pool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "stub"})
	}
}

func listInvoices(pool *pgxpool.Pool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON([]any{})
	}
}

func midtransCheckout(cfg config.Config, pool *pgxpool.Pool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"token": "snap_token_stub", "redirect_url": "https://app.sandbox.midtrans.com/snap/v3/redirection/mock"})
	}
}

func midtransWebhook(cfg config.Config, pool *pgxpool.Pool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	}
}
