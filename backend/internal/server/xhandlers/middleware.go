package xhandlers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

func RequestID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Get("X-Request-ID")
		if id == "" {
			id = utils.UUID()
			c.Set("X-Request-ID", id)
		}
		c.Locals("request_id", id)
		return c.Next()
	}
}

func Logger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()
		tenant, _ := c.Locals("tenant_slug").(string)
		log.Info().
			Str("method", c.Method()).
			Str("path", c.Path()).
			Int("status", c.Response().StatusCode()).
			Str("tenant", tenant).
			Msg("request")
		return err
	}
}

func TenantResolver(appDomain string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		host := c.Hostname()
		parts := strings.Split(host, ".")
		if len(parts) > 2 {
			c.Locals("tenant_slug", parts[0])
		} else {
			if h := c.Get("X-Tenant"); h != "" {
				c.Locals("tenant_slug", h)
			}
		}
		return c.Next()
	}
}

func WithTenant(pool *pgxpool.Pool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		slug, _ := c.Locals("tenant_slug").(string)
		if slug == "" {
			return fiber.NewError(fiber.StatusBadRequest, "missing tenant")
		}
		_, err := pool.Exec(c.Context(), `set local search_path = "`+slug+`", public`)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "invalid tenant")
		}
		return c.Next()
	}
}

func WithPublic(pool *pgxpool.Pool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		_, err := pool.Exec(c.Context(), `set local search_path = public`)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "db")
		}
		return c.Next()
	}
}
