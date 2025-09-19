Frontend (Nuxt 3) – Sprint 1 Skeleton

Prerequisites
- Node.js 20 LTS (use nvm)
- pnpm (installed globally)

Quick start
1) Use Node 20
   nvm install 20
   nvm use 20

2) Install pnpm (if not present)
   npm i -g pnpm

3) Install deps
   pnpm install

4) Run dev
   pnpm dev
   App: http://localhost:3000

5) Build for production
   pnpm build
   pnpm preview

Environment
- Copy .env.example to .env and set:
  - NUXT_PUBLIC_API_BASE (default http://localhost:8080/api)
  - NUXT_PUBLIC_TENANT (dev convenience header, e.g. demo)
  - NUXT_PUBLIC_MIDTRANS_CLIENT_KEY (Sandbox client key)

Notes
- Use X-Tenant header in dev (plugin http.ts sets it from NUXT_PUBLIC_TENANT).
- Do not commit .env or any secrets.
