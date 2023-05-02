<template>
   <h2>Master File Audit</h2>
   <div  v-if="auditStore.loading" class="wait-wrap">
      <WaitSpinner/>
   </div>
   <div v-else class="report">
      <BarChart :chartData="auditStore" :options="options"/>
      <div class="total">
         <label>Total Audited:</label><span class="total">{{auditStore.totalAudited}}</span>
      </div>
      <p class="error" v-if="auditStore.error">{{auditStore.error}}</p>
   </div>
</template>

<script setup>
import { onMounted, ref } from 'vue'
import { useAuditStore } from '@/stores/audit'
import WaitSpinner from '@/components/WaitSpinner.vue'
import { BarChart } from 'vue-chart-3'
import { Chart, registerables } from "chart.js"

Chart.register(...registerables)

const auditStore = useAuditStore()

const options = ref({
   responsive: true,
   title: {
      display: false,
   },
   legend: {
      display: false
   },
   plugins: {
      legend: {
         display: false,
      },
      colors: {
         enabled: false
      }
   },
})

onMounted( () => {
   auditStore.getAuditReport()
})
</script>

<style lang="scss" scoped>
.wait-wrap {
   text-align: center;
   margin-top: 10%;
}
.report {
   margin: 10px 50px;
   .total {
      text-align: center;
      margin: 20px 0;
   }
}
</style>
