<template>
   <div class="report-config">
      <label>Workflow:
         <select v-model="reportStore.workflowID">
            <option v-for="w in systemStore.workflows" :value="w.id" :key="`wf${w.id}`">{{w.name}}</option>
         </select>
      </label>
      <label>Start:
         <datepicker :typeable="true" v-model="reportStore.startDate" />
      </label>
      <label>End:
         <datepicker :typeable="true" v-model="reportStore.endDate" />
      </label>
   </div>
</template>

<script setup>
import {useSystemStore} from '@/stores/system'
import {useReportStore} from '@/stores/reporting'
import Datepicker from 'vue3-datepicker'

const systemStore = useSystemStore()
const reportStore = useReportStore()
</script>

<style lang="scss" scoped>
.report-config {
   display: flex;
   flex-flow: row nowrap;
   font-size: 0.9em;
   label {
      margin-right: 15px;
      font-weight: bold;
      color: var(--uvalib-text);
      :deep(input) {
         width: 85px;
      }
   }
   div {
      display: inline-block;
   }
}
</style>