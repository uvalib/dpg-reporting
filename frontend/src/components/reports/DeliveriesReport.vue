<template>
   <div class="reports-card">
      <h3>Patron Deliveries</h3>
      <div  v-if="reportStore.deliveries.loading" class="wait-wrap">
         <WaitSpinner/>
      </div>
      <div v-else class="report">
         <DoughnutChart :chartData="testData" />
         <div class="controls">
            <button @click="loadStats">Refresh Report</button>
         </div>
      </div>
   </div>
</template>

<script setup>
import {useReportStore} from '@/stores/reporting'
import WaitSpinner from '@/components/WaitSpinner.vue'
import { DoughnutChart } from 'vue-chart-3'
import { Chart, registerables } from "chart.js"

Chart.register(...registerables)

const testData = {
   labels: ['Paris', 'Nîmes', 'Toulon', 'Perpignan', 'Autre'],
   datasets: [
      {
         data: [30, 40, 60, 70, 5],
         backgroundColor: ['#77CEFF', '#0079AF', '#123E6B', '#97B0C4', '#A5C8ED'],
      },
   ],
}

const reportStore = useReportStore()

function loadStats() {
   reportStore.getDeliveriesReport("2021")
}
</script>

<style lang="scss" scoped>
.reports-card {
   margin: 10px;
   text-align: left;
   border: 1px solid var(--uvalib-grey-light);
   box-shadow: var(--box-shadow-light);
   h3 {
      background: var(--uvalib-grey-lightest);
      font-size: 1em;
      text-align: left;
      margin:0;
      padding: 5px 10px;
      border-bottom: 1px solid var(--uvalib-grey-light);
      font-weight: 500;;
   }
   .wait-wrap {
      text-align: center;
      padding: 30px 0 ;
   }
   .report {
      padding: 10px;
      .controls {
         border-top: 1px solid var(--uvalib-grey-lightest);
         display: flex;
         flex-flow: row wrap;
         justify-content: space-between;
         padding-top: 15px;
         margin-top: 5px;
         span {
            font-style: italic;
         }
         button {
            margin-left: auto;
         }
      }
   }
}
</style>

