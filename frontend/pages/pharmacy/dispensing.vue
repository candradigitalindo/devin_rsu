<template>
  <div>
    <h2>Dispensing</h2>
    <form @submit.prevent="submit">
      <input v-model="visitId" placeholder="Visit ID" />
      <input v-model="itemId" placeholder="Item ID" />
      <input v-model.number="qty" type="number" step="1" placeholder="Qty" />
      <button type="submit">Proses</button>
    </form>
    <div v-if="ok" class="ok">OK</div>
  </div>
</template>

<script setup lang="ts">
const { $api } = useNuxtApp() as any
const visitId = ref("")
const itemId = ref("")
const qty = ref<number>(1)
const ok = ref(false)
async function submit(){
  await $api('/tenant/pharmacy/dispensing',{ method:'POST', body:{ visit_id: visitId.value, item_id: itemId.value, qty: qty.value }})
  ok.value = true
}
</script>

<style>
.ok{margin-top:8px;color:#10b981}
</style>
