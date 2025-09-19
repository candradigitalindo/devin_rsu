import { ofetch } from "ofetch"

export default defineNuxtPlugin(() => {
  const config = useRuntimeConfig()
  const base = config.public.apiBaseUrl as string
  const api = ofetch.create({
    baseURL: base,
    onRequest({ options }) {
      const host = window.location.hostname
      const parts = host.split(".")
      const slug = parts.length > 2 ? parts[0] : "demo"
      options.headers = options.headers || {}
      ;(options.headers as Record<string, string>)["X-Tenant"] = slug
    }
  })
  return {
    provide: { api }
  }
})
