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
  - go mod tidy, run cmd/api/main.go
- Frontend
  - Env: see frontend/.env.example
  - pnpm install, pnpm dev
- Database
  - Apply docs/saas-rs-db-ddl.sql global section or backend/migrations/global/*.sql

Notes
- Tenant subdomain parsing in dev can use fixed X-Tenant header
- Do not commit secrets
