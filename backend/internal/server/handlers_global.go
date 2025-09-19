package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func listPlans(pool *pgxpool.Pool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		rows, err := pool.Query(c.Context(), `select id, code, name, features, price_rules from plans order by code`)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "db")
		}
		defer rows.Close()
		type Plan struct {
			ID         string      `json:"id"`
			Code       string      `json:"code"`
			Name       string      `json:"name"`
			Features   interface{} `json:"features"`
			PriceRules interface{} `json:"price_rules"`
		}
		var out []Plan
		for rows.Next() {
			var p Plan
			if err := rows.Scan(&p.ID, &p.Code, &p.Name, &p.Features, &p.PriceRules); err != nil {
				return fiber.NewError(fiber.StatusInternalServerError, "db")
			}
			out = append(out, p)
		}
		return c.JSON(out)
	}
}
