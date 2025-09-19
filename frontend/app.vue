<template>
  <div class="app">
    <header class="header">
      <h1>SaaS Klinik/RS</h1>
      <div class="tenant">{{ tenantSlug }}</div>
    </header>
    <nav class="nav">
      <NuxtLink to="/dashboard">Dashboard</NuxtLink>
      <NuxtLink to="/pharmacy">Farmasi</NuxtLink>
      <NuxtLink to="/billing">Billing</NuxtLink>
      <NuxtLink to="/reports/expiry">Expiry</NuxtLink>
    </nav>
    <main>
      <NuxtPage />
    </main>
  </div>
</template>

<script setup lang="ts">
const tenantSlug = computed(() => {
  const h = window.location.hostname.split(".")
  if (h.length > 2) return h[0]
  return "demo"
})
</script>

<style>
.app { font-family: Inter, system-ui, Arial, sans-serif; }
.header { display: flex; justify-content: space-between; padding: 12px 16px; border-bottom: 1px solid #eee; }
.nav { display: flex; gap: 12px; padding: 8px 16px; border-bottom: 1px solid #f0f0f0; }
</style>
