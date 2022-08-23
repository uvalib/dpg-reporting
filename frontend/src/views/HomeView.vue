<template>
   <main>
      <h2>DPG Statistics</h2>
       <div class="date-range">
         <label>Statistics from: </label>
         <select v-model="statsStore.dateRangeType">
            <option value="before">BEFORE</option>
            <option value="after">AFTER</option>
            <option value="between">BETWEEN</option>
         </select>

         <datepicker :typeable="true" v-model="statsStore.startDate" />
         <template v-if="statsStore.dateRangeType == 'between' ">
            <span class="date-sep">and</span>
            <datepicker :typeable="true" :clearable="true" v-model="statsStore.endDate" />
         </template>

         <button class="all-btn" @click="getAllClicked">Get All Statistics</button>
      </div>
      <div class="stats">
         <div class="column">
            <StorageStats />
            <ImageStats />
            <ArchiveStats />
         </div>
         <div class="column">
             <MetadataStats />
         </div>
      </div>
      <div class="stats">
         <div class="column">
             <h3>Recent Virgo Publications</h3>
             <div  v-if="statsStore.publishedStats.loading" class="wait-wrap">
               <WaitSpinner/>
            </div>
            <div v-else class="scroller">
               <table>
                  <tr>
                     <th></th><th>Title</th><th>Thumbnail</th><th>Details</th>
                  </tr>
                  <tr v-for="(rec,idx) in statsStore.publishedStats.virgo" :key="`as${rec.id}`">
                     <td class="num">{{idx+1}}.</td>
                     <td class="title">{{rec.title}}</td>
                     <td><img :src="rec.thumbURL"/></td>
                     <td><a :href="rec.adminURL" target="_blank">Details</a></td>
                  </tr>
               </table>
            </div>
         </div>
         <div class="column">
             <h3>Recent ArchivesSpace Publications</h3>
             <div  v-if="statsStore.publishedStats.loading" class="wait-wrap">
               <WaitSpinner/>
            </div>
            <div v-else class="scroller">
               <table>
                  <tr>
                     <th></th><th>Title</th><th>Details</th><th>Link</th>
                  </tr>
                  <tr v-for="(rec,idx) in statsStore.publishedStats.archivesSpace" :key="`as${rec.id}`">
                     <td class="num">{{idx+1}}.</td>
                     <td class="title">{{rec.title}}</td>
                     <td><a :href="rec.adminURL" target="_blank">Details</a></td>
                     <td><a :href="rec.externalURL" target="_blank">ArchivesSpace</a></td>
                  </tr>
               </table>
            </div>
         </div>
      </div>
   </main>
</template>

<script setup>
import Datepicker from 'vue3-datepicker'
import {useStatsStore} from '@/stores/statistics'
import ImageStats from '../components/ImageStats.vue'
import StorageStats from '../components/StorageStats.vue';
import MetadataStats from '../components/MetadataStats.vue';
import { onMounted } from 'vue'
import ArchiveStats from '../components/ArchiveStats.vue';
import WaitSpinner from '../components/WaitSpinner.vue';

const statsStore = useStatsStore()

onMounted( () => {
   statsStore.getAllStats(false)
})

function getAllClicked() {
   statsStore.getAllStats(true)
}
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
.wait-wrap {
   padding: 20px 10px;
}
.stats {
   margin: 10px 50px;
   display: flex;
   flex-flow: row wrap;
   .column {
      width: 48%;
   }
   h3 {
      margin: 10px 0 5px 0;
   }
   .scroller {
      text-align: left;
      font-size: 0.85em;
      margin: 0 20px 50px 0px;
      table {
         border-collapse: collapse;
         border: 1px solid #dedede;
         box-shadow: var(--box-shadow-light);
         th {
            background-color: #efefef;
            text-align: left;
            padding: 4px 10px 4px 5px;
            border-bottom: 1px solid #ccc;
         }
         td {
            vertical-align: middle;
            background: white;
            border-bottom: 1px solid #dedede;
            padding: 4px 10px 4px 5px;
         }
      }
   }
}
</style>
