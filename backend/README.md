Backend (Go Fiber) – Sprint 1 Skeleton

Prerequisites
- Go 1.22+ (locally available at /usr/local/go/bin/go on this VM)
- PostgreSQL running locally
- DATABASE_URL configured

Setup
1) cp .env.example .env
2) Edit .env to set DATABASE_URL and optional configs

Global Migrations + Seed
- Apply global schema and seed plans:
  /usr/local/go/bin/go run ./cmd/migrate

Run API
- Build:
  /usr/local/go/bin/go build ./cmd/api
- Run:
  ./api

Key Endpoints (prefix: /api)
- Health: GET /healthz
- Plans (public): GET /global/plans
- Admin: POST /admin/tenants  { slug, name, plan_id }
  - Creates schema and applies tenant migrations

Tenant Endpoints (require tenant via subdomain or header X-Tenant)
- Categories: GET/POST /tenant/pharmacy/categories
- Items: GET/POST /tenant/pharmacy/items
- Batches: GET/POST /tenant/pharmacy/items/:id/batches
- Dispensing (stub): POST /tenant/pharmacy/dispensing
- Expiry Report: GET /tenant/reports/pharmacy/expiry?before=YYYY-MM-DD

Billing (public scope for MVP)
- Invoices: POST/GET /billing/invoices
- Midtrans: POST /billing/midtrans/checkout, POST /billing/midtrans/webhook

Notes
- In dev, you can pass X-Tenant: <slug> header to target a tenant.
- The tenant provisioning reads and applies backend/migrations/tenant/001_init.sql after creating the schema.
- Do not commit secrets. Use .env for local only.
