<template>
   <div class="stats-card">
      <h3>Storage Statistics</h3>
      <div  v-if="reportStore.storageStats.loading" class="wait-wrap">
         <WaitSpinner/>
      </div>
      <div v-else class="stats">
         <p v-if="reportStore.storageStats.rangeText">Images created {{reportStore.storageStats.rangeText}}</p>
         <dl>
            <dt>Total Image Size (GB):</dt>
            <dd>{{numberWithCommas(reportStore.storageStats.all)}}</dd>
            <dt>DL Total Image Size (GB):</dt>
            <dd>{{numberWithCommas(reportStore.storageStats.DL)}}</dd>
         </dl>
         <div class="controls">
            <button @click="loadStats">Refresh Stats</button>
         </div>
      </div>
   </div>
</template>

<script setup>
import {useReportStore} from '@/stores/reporting'
import WaitSpinner from './WaitSpinner.vue';
const reportStore = useReportStore()

function loadStats() {
   reportStore.getStorageSats()
}
function numberWithCommas(num) {
    return num.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",");
}
</script>

<style lang="scss" scoped>
.stats-card {
   margin: 10px;
   text-align: left;
   max-width: 40%;
   min-width: 300px;
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
   .stats {
      padding: 10px;
      p {
         font-size: 0.9em;
         padding: 0;
         margin: 0 5px;
      }
      .controls {
         text-align: right;
      }

      dl {
      margin: 10px 30px 0 30px;
      display: inline-grid;
      grid-template-columns: max-content 2fr;
      grid-column-gap: 10px;
      font-size: 0.9em;
      text-align: left;
      box-sizing: border-box;

      dt {
         font-weight: bold;
         text-align: right;
      }
      dd {
         margin: 0 0 10px 0;
         word-break: break-word;
         -webkit-hyphens: auto;
         -moz-hyphens: auto;
         hyphens: auto;
      }
   }
   }
}
</style>

