<template>
   <main>
      <h2>Digitization Reports</h2>
      <div class="control-bar">
         <ReportConfig />
         <button @click="loadStats">Generate Reports</button>
      </div>
      <div class="reports">
         <div class="column">
            <PageTimeReport />
            <ProductivityReport />
         </div>
         <div class="column">
            <ProblemsReport />
            <RejectionsReport />
            <RatesReport />
         </div>
         <DeliveriesReport />
      </div>
   </main>
</template>

<script setup>
import DeliveriesReport from '@/components/reports/DeliveriesReport.vue'
import ProductivityReport from '@/components/reports/ProductivityReport.vue'
import ReportConfig from '../components/reports/ReportConfig.vue'
import {useReportStore} from '@/stores/reporting'
import ProblemsReport from '../components/reports/ProblemsReport.vue'
import PageTimeReport from '../components/reports/PageTimeReport.vue'
import RejectionsReport from '../components/reports/RejectionsReport.vue'
import RatesReport from '../components/reports/RatesReport.vue'

const reportStore = useReportStore()

function loadStats() {
   reportStore.getProductivityReport(reportStore.workflowID, reportStore.startDate, reportStore.endDate)
   reportStore.getProblemsReport(reportStore.workflowID, reportStore.startDate, reportStore.endDate)
   reportStore.getPageTimesReport(reportStore.workflowID, reportStore.startDate, reportStore.endDate)
   reportStore.getRejectionsReport(reportStore.workflowID, reportStore.startDate, reportStore.endDate)
   reportStore.getRatesReport(reportStore.workflowID, reportStore.startDate, reportStore.endDate)
}
</script>

<style lang="scss" scoped>
.control-bar {
   display: flex;
   flex-flow: row nowrap;
   justify-content: space-between;
   align-items: center;
   padding: 20px;
   border-bottom: 1px solid var(--uvalib-grey-light);
   margin-bottom: 15px;
   border-top: 1px solid var(--uvalib-grey-light);
}
.reports {
   margin: 10px 50px;
   display: flex;
   flex-flow: row wrap;
   .column {
      width: 48%;
   }
}
</style>
