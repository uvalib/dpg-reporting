<template>
   <main>
      <h2>Digitization Reports</h2>
      <div class="reports">
         <div class="column">
            <DeliveriesReport />
         </div>
      </div>
   </main>
</template>

<script setup>
import DeliveriesReport from '@/components/reports/DeliveriesReport.vue'
</script>

<style lang="scss" scoped>
.reports {
   margin: 10px 50px;
   display: flex;
   flex-flow: row wrap;
   .column {
      width: 48%;
   }
}
</style>
