<template>
   <main>
      <h2>DPG Statistics</h2>
       <div class="date-range">
         <label>Statistics from: </label>
         <select v-model="reportStore.dateRangeType">
            <option value="before">BEFORE</option>
            <option value="after">AFTER</option>
            <option value="between">BETWEEN</option>
         </select>

         <datepicker :typeable="true" :clearable="true" v-model="reportStore.startDate" />
         <template v-if="reportStore.dateRangeType == 'between' ">
            <span class="date-sep">and</span>
            <datepicker :typeable="true" :clearable="true" v-model="reportStore.endDate" />
         </template>

         <!-- <button class="all-btn" @click="getAllClicked">Get All Statistics</button> -->
      </div>
      <div class="stats">
         <ImageStats />
         <StorageStats />
         <MetadataStats />
      </div>
   </main>
</template>

<script setup>
import Datepicker from 'vue3-datepicker'
import {useReportStore} from '@/stores/reporting'
import ImageStats from '../components/ImageStats.vue'
import StorageStats from '../components/StorageStats.vue';
import MetadataStats from '../components/MetadataStats.vue';
const reportStore = useReportStore()

// function getAllClicked() {
//    alert("not implemented")
// }
</script>

<style lang="scss" scoped>
.date-range {
   display:flex;
   flex-flow: row nowrap;
   padding: 20px;
   border-bottom: 1px solid var(--uvalib-grey-light);
   margin-bottom: 35px;
   border-top: 1px solid var(--uvalib-grey-light);

   label {
      font-weight: bold;
      margin-right: 10px;
   }
   select {
      margin-right: 10px;
   }
   .date-sep {
      margin-right: 10px;
      display: inline-block;
   }
   .all-btn {
      margin-left: auto;
   }
}
.stats {
   margin: 10px 50px;
}
</style>
