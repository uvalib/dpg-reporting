<template>
   <div class="stats-card">
      <h3>Archive Statistics</h3>
      <div  v-if="statsStore.archiveStats.loading" class="wait-wrap">
         <WaitSpinner/>
      </div>
      <div v-else class="stats">
         <dl>
            <dt>Total Bound Volumes:</dt>
            <dd>{{numberWithCommas(statsStore.archiveStats.bound)}}</dd>
            <dt>Total MSS / Unbound Sheets:</dt>
            <dd>{{numberWithCommas(statsStore.archiveStats.manuscript)}}</dd>
            <dt>Total Photograph / AV Items:</dt>
            <dd>{{numberWithCommas(statsStore.archiveStats.photo)}}</dd>
         </dl>
         <div class="controls">
            <span v-if="statsStore.archiveStats.rangeText">Items archived {{statsStore.archiveStats.rangeText}}</span>
            <button @click="loadStats">Refresh Stats</button>
         </div>
      </div>
   </div>
</template>

<script setup>
import {useStatsStore} from '@/stores/statistics'
import WaitSpinner from './WaitSpinner.vue';
const statsStore = useStatsStore()

function loadStats() {
   statsStore.getArchiveSats()
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
      p {
         font-size: 0.9em;
         padding: 0;
         margin: 0 5px;
      }
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

