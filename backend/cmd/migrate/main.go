package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/candra/saas-rs-backend/internal/config"
	"github.com/candra/saas-rs-backend/internal/db"
)

func main() {
	cfg := config.Load()
	if cfg.DatabaseURL == "" {
		log.Fatal("DATABASE_URL is required")
	}
	pool, err := db.NewPool(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	ctx := context.Background()

	_, err = pool.Exec(ctx, `
CREATE EXTENSION IF NOT EXISTS pgcrypto;
DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'ulid_char') THEN
    CREATE DOMAIN ulid_char AS CHAR(26)
      CHECK (VALUE ~ '^[0-9A-HJKMNP-TV-Z]{26}$');
  END IF;
END$$;

CREATE TABLE IF NOT EXISTS plans (
  id            ulid_char PRIMARY KEY,
  code          TEXT NOT NULL CHECK (code IN ('KLINIK','RS')),
  name          TEXT NOT NULL,
  features      JSONB NOT NULL DEFAULT '{}'::jsonb,
  price_rules   JSONB NOT NULL DEFAULT '{}'::jsonb,
  created_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at    TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS tenants (
  id            ulid_char PRIMARY KEY,
  slug          TEXT UNIQUE NOT NULL,
  name          TEXT NOT NULL,
  status        TEXT NOT NULL DEFAULT 'active' CHECK (status IN ('active','suspended','pending')),
  plan_id       ulid_char NOT NULL REFERENCES plans(id),
  created_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at    TIMESTAMPTZ NOT NULL DEFAULT now()
);
`)
	if err != nil {
		log.Fatalf("global migration failed: %v", err)
	}

	_, err = pool.Exec(ctx, `
-- seed plans if not present
INSERT INTO plans (id, code, name, features, price_rules)
SELECT '01H00000000000000000000000', 'KLINIK', 'Klinik', '{}'::jsonb, '{}'::jsonb
WHERE NOT EXISTS (SELECT 1 FROM plans WHERE code='KLINIK');

INSERT INTO plans (id, code, name, features, price_rules)
SELECT '01H00000000000000000000001', 'RS', 'Rumah Sakit', '{}'::jsonb, '{}'::jsonb
WHERE NOT EXISTS (SELECT 1 FROM plans WHERE code='RS');
`)
	if err != nil {
		log.Fatalf("plan seed failed: %v", err)
	}

	fmt.Fprintln(os.Stdout, "Global migrations applied and plans seeded.")
}
