package server

import (
	"crypto/rand"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/candra/saas-rs-backend/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/oklog/ulid/v2"
)
func newULID() string {
	id := ulid.MustNew(ulid.Timestamp(time.Now()), rand.Reader)
	return id.String()
}
func readTenantMigration() ([]byte, error) {
	paths := []string{
		"migrations/tenant/001_init.sql",
		"../migrations/tenant/001_init.sql",
		"../../migrations/tenant/001_init.sql",
		"backend/migrations/tenant/001_init.sql",
	}
	for _, p := range paths {
		if _, err := os.Stat(p); err == nil {
			return ioutil.ReadFile(p)
		}
	}
	return nil, os.ErrNotExist
}



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
		id := newULID()
		_, err := pool.Exec(c.Context(), `insert into tenants(id,slug,name,plan_id,status) values ($1, $2, $3, $4, 'active') on conflict (slug) do nothing`, id, req.Slug, req.Name, req.Plan)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "tenant invalid")
		}
		_, err = pool.Exec(c.Context(), `create schema if not exists "`+req.Slug+`"`)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "schema")
		}
		sqlBytes, err := readTenantMigration()
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "migration read")
		}
		tx, err := pool.Begin(c.Context())
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "migration begin")
		}
		defer tx.Rollback(c.Context())
		if _, err = tx.Exec(c.Context(), `set local search_path = "`+req.Slug+`", public`); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "migration path")
		}
		stmts := strings.Split(string(sqlBytes), ";")
		for _, s := range stmts {
			s2 := strings.TrimSpace(s)
			if s2 == "" {
				continue
			}
			if _, err = tx.Exec(c.Context(), s2); err != nil {
				return fiber.NewError(fiber.StatusInternalServerError, "migration apply")
			}
		}
		if err = tx.Commit(c.Context()); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "migration commit")
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
		id := newULID()
		_, err := pool.Exec(c.Context(), `insert into pharmacy_categories(id,name,code) values ($1, $2, $3)`, id, req.Name, req.Code)
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
		id := newULID()
		_, err := pool.Exec(c.Context(), `insert into pharmacy_items(id,category_id,name,sku,uom,min_stock) values ($1, $2, $3, $4, $5, coalesce($6,0))`,
			id, req.CategoryID, req.Name, req.SKU, req.UOM, req.MinStock)
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
		id := newULID()
		_, err := pool.Exec(c.Context(), `insert into pharmacy_item_batches(id,item_id,batch_no,expiry_date,qty_on_hand) values ($1, $2, $3, $4, $5)`,
			id, itemID, req.BatchNo, req.ExpiryDate, req.Qty)
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
