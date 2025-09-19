<template>
  <div>
    <h2>Dashboard</h2>
    <div class="cards">
      <div class="card">Antrean</div>
      <div class="card">Kunjungan</div>
      <div class="card">Penjualan</div>
      <div class="card">Stok Rendah</div>
    </div>
  </div>
</template>
<style>
.cards{display:grid;grid-template-columns:repeat(4,1fr);gap:12px}
.card{padding:16px;border:1px solid #eee;border-radius:8px}
</style>
