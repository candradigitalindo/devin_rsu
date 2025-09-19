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
