package server

import (
	"io/ioutil"
	"strings"
	"time"

	"github.com/candra/saas-rs-backend/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type createTenantReq struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
	Plan string `json:"plan_id"`
}

func createTenant(cfg config.Config, pool *pgxpool.Pool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req createTenantReq
		if err := c.BodyParser(&req); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "invalid")
		}
		req.Slug = strings.ToLower(req.Slug)
		_, err := pool.Exec(c.Context(), `insert into tenants(id,slug,name,plan_id,status) values (left(replace(gen_random_uuid()::text,'-',''),26), $1, $2, $3, 'active') on conflict (slug) do nothing`, req.Slug, req.Name, req.Plan)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "tenant invalid")
		}
		_, err = pool.Exec(c.Context(), `create schema if not exists "`+req.Slug+`"`)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "schema")
		}
		sqlBytes, err := ioutil.ReadFile("backend/migrations/tenant/001_init.sql")
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "migration read")
		}
		_, err = pool.Exec(c.Context(), `set local search_path = "`+req.Slug+`", public; `+string(sqlBytes))
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "migration apply")
		}
		return c.Status(201).JSON(fiber.Map{"slug": req.Slug, "status": "provisioned"})
	}
}

func listCategories(pool *pgxpool.Pool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		rows, err := pool.Query(c.Context(), `select id, name, code, parent_id from pharmacy_categories order by name`)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "db")
		}
		defer rows.Close()
		type Cat struct {
			ID       string `json:"id"`
			Name     string `json:"name"`
			Code     string `json:"code"`
			ParentID string `json:"parent_id"`
		}
		var out []Cat
		for rows.Next() {
			var x Cat
			if err := rows.Scan(&x.ID, &x.Name, &x.Code, &x.ParentID); err != nil {
				return fiber.NewError(fiber.StatusInternalServerError, "db")
			}
			out = append(out, x)
		}
		return c.JSON(out)
	}
}

func createCategory(pool *pgxpool.Pool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		type Req struct {
			Name string `json:"name"`
			Code string `json:"code"`
		}
		var req Req
		if err := c.BodyParser(&req); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "invalid")
		}
		_, err := pool.Exec(c.Context(), `insert into pharmacy_categories(id,name,code) values (left(replace(gen_random_uuid()::text,'-',''),26), $1, $2)`, req.Name, req.Code)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "db")
		}
		return c.SendStatus(201)
	}
}

func listItems(pool *pgxpool.Pool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		rows, err := pool.Query(c.Context(), `select id, category_id, name, sku, uom, min_stock, is_active from pharmacy_items order by name`)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "db")
		}
		defer rows.Close()
		type Item struct {
			ID         string  `json:"id"`
			CategoryID *string `json:"category_id"`
			Name       string  `json:"name"`
			SKU        *string `json:"sku"`
			UOM        string  `json:"uom"`
			MinStock   float64 `json:"min_stock"`
			IsActive   bool    `json:"is_active"`
		}
		var out []Item
		for rows.Next() {
			var x Item
			if err := rows.Scan(&x.ID, &x.CategoryID, &x.Name, &x.SKU, &x.UOM, &x.MinStock, &x.IsActive); err != nil {
				return fiber.NewError(fiber.StatusInternalServerError, "db")
			}
			out = append(out, x)
		}
		return c.JSON(out)
	}
}

func createItem(pool *pgxpool.Pool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		type Req struct {
			CategoryID *string `json:"category_id"`
			Name       string  `json:"name"`
			SKU        *string `json:"sku"`
			UOM        string  `json:"uom"`
			MinStock   *float64 `json:"min_stock"`
		}
		var req Req
		if err := c.BodyParser(&req); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "invalid")
		}
		_, err := pool.Exec(c.Context(), `insert into pharmacy_items(id,category_id,name,sku,uom,min_stock) values (left(replace(gen_random_uuid()::text,'-',''),26), $1, $2, $3, $4, coalesce($5,0))`,
			req.CategoryID, req.Name, req.SKU, req.UOM, req.MinStock)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "db")
		}
		return c.SendStatus(201)
	}
}

func listBatches(pool *pgxpool.Pool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		itemID := c.Params("id")
		rows, err := pool.Query(c.Context(), `select id, item_id, batch_no, expiry_date, qty_on_hand, received_at from pharmacy_item_batches where item_id=$1 order by expiry_date asc`, itemID)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "db")
		}
		defer rows.Close()
		type Batch struct {
			ID         string    `json:"id"`
			ItemID     string    `json:"item_id"`
			BatchNo    string    `json:"batch_no"`
			ExpiryDate time.Time `json:"expiry_date"`
			QtyOnHand  float64   `json:"qty_on_hand"`
			ReceivedAt time.Time `json:"received_at"`
		}
		var out []Batch
		for rows.Next() {
			var x Batch
			if err := rows.Scan(&x.ID, &x.ItemID, &x.BatchNo, &x.ExpiryDate, &x.QtyOnHand, &x.ReceivedAt); err != nil {
				return fiber.NewError(fiber.StatusInternalServerError, "db")
			}
			out = append(out, x)
		}
		return c.JSON(out)
	}
}

func createBatch(pool *pgxpool.Pool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		itemID := c.Params("id")
		type Req struct {
			BatchNo   string    `json:"batch_no"`
			ExpiryDate string   `json:"expiry_date"`
			Qty       float64   `json:"qty"`
		}
		var req Req
		if err := c.BodyParser(&req); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "invalid")
		}
		_, err := pool.Exec(c.Context(), `insert into pharmacy_item_batches(id,item_id,batch_no,expiry_date,qty_on_hand) values (left(replace(gen_random_uuid()::text,'-',''),26), $1, $2, $3, $4)`,
			itemID, req.BatchNo, req.ExpiryDate, req.Qty)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "db")
		}
		return c.SendStatus(201)
	}
}

func dispensing(pool *pgxpool.Pool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendStatus(201)
	}
}

func reportExpiry(pool *pgxpool.Pool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		q := c.Query("before")
		rows, err := pool.Query(c.Context(), `select id, item_id, batch_no, expiry_date, qty_on_hand, received_at from pharmacy_item_batches where expiry_date < $1 order by expiry_date asc`, q)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "db")
		}
		defer rows.Close()
		type Row struct {
			ID         string    `json:"id"`
			ItemID     string    `json:"item_id"`
			BatchNo    string    `json:"batch_no"`
			ExpiryDate time.Time `json:"expiry_date"`
			QtyOnHand  float64   `json:"qty_on_hand"`
			ReceivedAt time.Time `json:"received_at"`
		}
		var out []Row
		for rows.Next() {
			var x Row
			if err := rows.Scan(&x.ID, &x.ItemID, &x.BatchNo, &x.ExpiryDate, &x.QtyOnHand, &x.ReceivedAt); err != nil {
				return fiber.NewError(fiber.StatusInternalServerError, "db")
			}
			out = append(out, x)
		}
		return c.JSON(out)
	}
}
