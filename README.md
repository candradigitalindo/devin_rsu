SaaS Klinik/RS – Skeleton (Sprint 1)

Stack
- Backend: Go (Fiber), PostgreSQL per-schema, ULID
- Frontend: Nuxt 3 (Vue)
- Payments: Midtrans Snap (sandbox)
- OpenAPI: docs/saas-rs-openapi.yaml
- DDL: docs/saas-rs-db-ddl.sql

Structure
- backend/
- frontend/
- docs/

Local Dev
- Backend
  - Env: see backend/.env.example
  - /usr/local/go/bin/go mod tidy
  - /usr/local/go/bin/go run ./cmd/migrate
  - /usr/local/go/bin/go run ./cmd/api
- Frontend
  - Use Node 20 LTS (nvm use 20). See frontend/README.md
  - Env: see frontend/.env.example
  - pnpm install, pnpm dev
- Database
  - Apply global migration via cmd/migrate or see backend/migrations/global/*.sql

Notes
- In dev, set X-Tenant header via frontend plugin or API client
- Do not commit secrets (.env). .nvmrc pins Node 20 for FE
