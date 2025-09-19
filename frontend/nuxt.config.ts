export default defineNuxtConfig({
  ssr: true,
  modules: [],
  runtimeConfig: {
    public: {
      apiBaseUrl: process.env.NUXT_PUBLIC_API_BASE_URL || "http://localhost:8080/api",
      authIssuer: process.env.NUXT_PUBLIC_AUTH_ISSUER || "",
      authClientId: process.env.NUXT_PUBLIC_AUTH_CLIENT_ID || ""
    }
  }
  ,
  nitro: {
    compatibilityDate: '2025-09-18'
  }
})
