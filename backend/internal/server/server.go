package server

import (
	"github.com/candra/saas-rs-backend/internal/config"
	"github.com/candra/saas-rs-backend/internal/server/xhandlers"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func New(cfg config.Config, pool *pgxpool.Pool) *fiber.App {
	app := fiber.New()
	app.Use(xhandlers.RequestID())
	app.Use(xhandlers.Logger())

	api := app.Group("/api")

	api.Get("/healthz", func(c *fiber.Ctx) error { return c.SendStatus(200) })
	api.Get("/readyz", func(c *fiber.Ctx) error { return c.SendStatus(200) })

	global := api.Group("/global", xhandlers.WithPublic(pool))
	global.Get("/plans", listPlans(pool))

	tenant := api.Group("/tenant", xhandlers.TenantResolver(cfg.AppDomain), xhandlers.WithTenant(pool))
	tenant.Get("/pharmacy/categories", listCategories(pool))
	tenant.Post("/pharmacy/categories", createCategory(pool))
	tenant.Get("/pharmacy/items", listItems(pool))
	tenant.Post("/pharmacy/items", createItem(pool))
	tenant.Get("/pharmacy/items/:id/batches", listBatches(pool))
	tenant.Post("/pharmacy/items/:id/batches", createBatch(pool))
	tenant.Post("/pharmacy/dispensing", dispensing(pool))
	tenant.Get("/reports/pharmacy/expiry", reportExpiry(pool))

	billing := api.Group("/billing", xhandlers.WithPublic(pool))
	billing.Post("/invoices", createInvoice(pool))
	billing.Get("/invoices", listInvoices(pool))
	billing.Post("/midtrans/checkout", midtransCheckout(cfg, pool))
	billing.Post("/midtrans/webhook", midtransWebhook(cfg, pool))

	admin := api.Group("/admin", xhandlers.WithPublic(pool))
	admin.Post("/tenants", createTenant(cfg, pool))

	return app
}
