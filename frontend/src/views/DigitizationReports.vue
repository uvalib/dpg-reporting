<template>
  <main>
    <h2>Digitization Reports</h2>
    <div class="reports">
      <DeliveriesReport />
    </div>
  </main>
</template>

<script setup>
import DeliveriesReport from '@/components/reports/DeliveriesReport.vue'
</script>

<style lang="scss" scoped>
.reports {
  margin: 10px 50px;
}
</style>
