CREATE TABLE IF NOT EXISTS pharmacy_categories (
  id              ulid_char PRIMARY KEY,
  parent_id       ulid_char,
  name            TEXT NOT NULL,
  code            TEXT,
  created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS pharmacy_items (
  id              ulid_char PRIMARY KEY,
  category_id     ulid_char REFERENCES pharmacy_categories(id),
  name            TEXT NOT NULL,
  sku             TEXT,
  uom             TEXT NOT NULL,
  min_stock       NUMERIC(18,3) DEFAULT 0,
  is_active       BOOLEAN NOT NULL DEFAULT true,
  created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS pharmacy_item_batches (
  id              ulid_char PRIMARY KEY,
  item_id         ulid_char NOT NULL REFERENCES pharmacy_items(id) ON DELETE CASCADE,
  batch_no        TEXT NOT NULL,
  expiry_date     DATE NOT NULL,
  qty_on_hand     NUMERIC(18,3) NOT NULL DEFAULT 0,
  received_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
  UNIQUE (item_id, batch_no, expiry_date)
);

CREATE INDEX IF NOT EXISTS idx_batches_expiry ON pharmacy_item_batches (expiry_date);
