<template>
  <div>
    <h2>Billing</h2>
    <button @click="createCheckout">Bayar Invoice (Stub)</button>
    <div v-if="token">Snap Token: {{ token }}</div>
  </div>
</template>

<script setup lang="ts">
const { $api } = useNuxtApp() as any
const token = ref("")
async function createCheckout(){
  const r = await $api('/billing/midtrans/checkout',{ method:'POST', body:{ invoice_id:"stub" } })
  token.value = r.token
  window.open(r.redirect_url, '_blank')
}
</script>
