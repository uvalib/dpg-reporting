<template>
   <div class="reports-card">
      <h3>Average Page Completion Time</h3>
      <div  v-if="reportStore.pageTimes.loading" class="wait-wrap">
         <WaitSpinner/>
      </div>
      <div v-else class="report">
         <BarChart :chartData="reportStore.pageTimes" :options="options"/>
         <table class="raw-data">
            <tbody>
               <tr>
                  <th>Category</th>
                  <th>Units</th>
                  <th>Total Mins</th>
                  <th>Total Pages</th>
                  <th>Avg. Mins</th>
               </tr>
               <tr v-for="(r,idx) in reportStore.pageTimes.raw" :key="`raw${idx}`">
                  <td>{{r.category}}</td>
                  <td>{{r.units}}</td>
                  <td>{{r.totalMins}}</td>
                  <td>{{r.totalPages}}</td>
                  <td>{{r.avgPageTime}}</td>
               </tr>
            </tbody>
         </table>
         <p class="error" v-if="reportStore.pageTimes.error">{{reportStore.pageTimes.error}}</p>
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
      scales:  {
         xAxis: {
            ticks: {
               autoSkip: false
            }
         }
      },
      plugins: {
        legend: {
            display: false,
        },
      },
    });

const reportStore = useReportStore()

onMounted( () => {
   if (reportStore.pageTimes.datasets.length == 0) {
      reportStore.getPageTimesReport(reportStore.workflowID, reportStore.startDate, reportStore.endDate)
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
   .raw-data {
      border-collapse: collapse;
      width: 100%;
      margin: 15px 15px;
      font-size: 0.9em;
      color: var(--uvalib-text);
      th {
         font-weight: normal;
         border-bottom: 1px solid var(--uvalib-grey-light);
      }
      th, td {
         padding: 5px 0;
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

