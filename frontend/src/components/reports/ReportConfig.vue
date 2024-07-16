<template>
   <div class="report-config">
      <label>Workflow:</label>
      <select v-model="reportStore.workflowID">
         <option v-for="w in systemStore.workflows" :value="w.id" :key="`wf${w.id}`">{{w.name}}</option>
      </select>
      <label>Start:</label>
      <VueDatePicker v-model="reportStore.startDate"  :enable-time-picker="false" format="yyyy-MM-dd" auto-apply :clearable="false" />

      <label>End:</label>
      <VueDatePicker v-model="reportStore.endDate"  :enable-time-picker="false" format="yyyy-MM-dd" auto-apply :clearable="false" />
   </div>
</template>

<script setup>
import {useSystemStore} from '@/stores/system'
import {useReportStore} from '@/stores/reporting'

const systemStore = useSystemStore()
const reportStore = useReportStore()
</script>

<style lang="scss" scoped>
.report-config {
   display: flex;
   flex-flow: row nowrap;
   align-items: center;
   font-size: 0.9em;
   gap: 10px;

   :deep(.dp__main) {
      max-width: 150px !important;
      input.dp__input {
         padding-left: 35px  !important;
      }
   }

   label {
      margin-left: 20px;
      font-weight: bold;
      color: var(--uvalib-text);
   }
}
</style>