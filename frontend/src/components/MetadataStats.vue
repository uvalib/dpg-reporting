<template>
   <div class="stats-card">
      <h3>Metadata Statistics</h3>
      <div  v-if="reportStore.metadataStats.loading" class="wait-wrap">
         <WaitSpinner/>
      </div>
      <div v-else class="stats">
         <dl>
            <dt>Metadata Count:</dt>
            <dd>{{numberWithCommas(reportStore.metadataStats.total)}}</dd>
            <dt>SIRSI Metadata Count:</dt>
            <dd>{{numberWithCommas(reportStore.metadataStats.sirsi)}}</dd>
            <dt>XML Metadata Count:</dt>
            <dd>{{numberWithCommas(reportStore.metadataStats.xml)}}</dd>
            <dt>DL Metadata Count:</dt>
            <dd>{{numberWithCommas(reportStore.metadataStats.totalDL)}}</dd>
            <dt>DL SIRSI Metadata Count:</dt>
            <dd>{{numberWithCommas(reportStore.metadataStats.sirsiDL)}}</dd>
            <dt>DL XML Metadata Count:</dt>
            <dd>{{numberWithCommas(reportStore.metadataStats.xmlDL)}}</dd>
             <dt>DPLA Metadata Count:</dt>
            <dd>{{numberWithCommas(reportStore.metadataStats.totalDPLA)}}</dd>
            <dt>DPLA SIRSI Metadata Count:</dt>
            <dd>{{numberWithCommas(reportStore.metadataStats.sirsiDPLA)}}</dd>
            <dt>DPLA XML Metadata Count:</dt>
            <dd>{{numberWithCommas(reportStore.metadataStats.xmlDPLA)}}</dd>
         </dl>
         <div class="controls">
            <span v-if="reportStore.metadataStats.rangeText">Metadata created {{reportStore.metadataStats.rangeText}}</span>
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
   reportStore.getMetadataSats()
}
function numberWithCommas(num) {
    return num.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",");
}
</script>

<style lang="scss" scoped>
.stats-card {
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
   .stats {
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

