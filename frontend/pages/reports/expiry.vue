<template>
  <div>
    <h2>Expiry Report</h2>
    <input type="date" v-model="before" />
    <button @click="load">Tampilkan</button>
    <ul>
      <li v-for="b in rows" :key="b.id">{{ b.item_name }} - {{ b.batch_no }} - {{ b.expiry_date }}</li>
    </ul>
  </div>
</template>

<script setup lang="ts">
const { $api } = useNuxtApp() as any
const before = ref<string>("")
const rows = ref<any[]>([])
async function load(){
  rows.value = await $api('/tenant/reports/pharmacy/expiry',{ params:{ before: before.value }})
}
</script>
