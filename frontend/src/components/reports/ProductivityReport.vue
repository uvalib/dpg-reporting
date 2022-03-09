<template>
   <div class="reports-card">
      <h3>Productivity</h3>
       <div  v-if="reportStore.productivity.loading" class="wait-wrap">
         <WaitSpinner/>
      </div>
      <div v-else class="report">
         <BarChart :chartData="reportStore.productivity" :options="options"/>
         <div class="total">
            <label>Total Completed Projects:</label><span class="total">{{reportStore.productivity.totalCompleted}}</span>
         </div>
         <p class="error" v-if="reportStore.productivity.error">{{reportStore.productivity.error}}</p>
      </div>
   </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import {useReportStore} from '@/stores/reporting'
import WaitSpinner from '@/components/WaitSpinner.vue'
import { BarChart } from 'vue-chart-3'
import { Chart, registerables } from "chart.js"

Chart.register(...registerables)

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
      },
    });

const reportStore = useReportStore()

onMounted( () => {
   if (reportStore.productivity.datasets.length == 0) {
      reportStore.getProductivityReport(reportStore.workflowID, reportStore.startDate, reportStore.endDate)
   }
})
</script>

<style lang="scss" scoped>
.reports-card {
   margin: 10px;
   text-align: left;
   border: 1px solid var(--uvalib-grey-light);
   box-shadow: var(--box-shadow-light);
   position: relative;
   min-height: 360px;
   h3 {
      background: var(--uvalib-grey-lightest);
      font-size: 1em;
      text-align: left;
      margin:0;
      padding: 5px 10px;
      border-bottom: 1px solid var(--uvalib-grey-light);
      font-weight: 500;;
   }
   div.total {
      margin: 10px 0 5px 0;
      text-align: center;
      span.total {
         margin-left: 10px;
      }
   }
   .wait-wrap {
      text-align: center;
      padding: 30px 0 ;
      position: absolute;
      top: 0;
      left: 0;
      right: 0;
      bottom: 0;
      background-color:rgba(255,255,255, 0.6);
      z-index: 1000;
      div.spinner {
         margin-top: 25%;
      }
   }
   .report {
      padding: 10px;
   }
}
</style>

