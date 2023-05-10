<template>
   <h2>Master File Audit</h2>
   <div  v-if="auditStore.loading" class="wait-wrap">
      <WaitSpinner/>
   </div>
   <div v-else class="report">
      <div class="control-bar">
         <label>Audit Year:</label>
         <select v-model="auditStore.targetYear">
            <option v-for="w in auditStore.auditYears" :value="w.value" :key="`wf${w.value}`">{{w.label}}</option>
         </select>
         <button @click="auditStore.getAuditReport()">Generate Audit Report</button>
      </div>
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
   .control-bar {
      display: flex;
      flex-flow: row nowrap;
      justify-content: flex-end;
      align-items: center;
      label, select {
         margin-right: 10px;
      }
   }
   .total {
      text-align: center;
      margin: 20px 0;
   }
}
</style>
