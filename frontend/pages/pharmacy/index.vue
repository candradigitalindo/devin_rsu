<template>
  <div>
    <h2>Farmasi - Stok</h2>
    <div class="filters">
      <select v-model="expired">
        <option :value="''">Semua</option>
        <option value="true">Expired</option>
        <option value="false">Non-Expired</option>
      </select>
      <input v-model="q" placeholder="Cari item" />
      <button @click="load">Terapkan</button>
    </div>
    <table>
      <thead>
        <tr><th>Nama</th><th>Kategori</th><th>Stok</th><th>Status</th></tr>
      </thead>
      <tbody>
        <tr v-for="it in items" :key="it.id">
          <td>{{ it.name }}</td>
          <td>{{ it.category_name || '-' }}</td>
          <td>{{ it.stock || 0 }}</td>
          <td><span :class="badgeClass(it)">{{ it.expiryStatus || '-' }}</span></td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup lang="ts">
const { $api } = useNuxtApp() as any
const items = ref<any[]>([])
const expired = ref<string>("")
const q = ref("")
function badgeClass(it:any){ return "badge" }
async function load(){
  const params:any = {}
  if (expired.value) params.expired = expired.value
  if (q.value) params.q = q.value
  items.value = await $api('/tenant/pharmacy/items',{ params })
}
onMounted(load)
</script>

<style>
.filters{display:flex;gap:8px;margin-bottom:12px}
.badge{padding:2px 8px;border-radius:12px;border:1px solid #ddd}
</style>
