<template>
   <header>
      <div class="main-header">
         <div class="library-link">
            <a target="_blank" href="https://library.virginia.edu">
               <UvaLibraryLogo />
            </a>
         </div>
         <div class="site-link">
            <router-link to="/">DPG Reporting</router-link>
            <p class="version">{{ systemStore.version }}</p>
         </div>
      </div>
   </header>
   <nav class="menu" role="menubar" aria-label="DPG Reporting Menu">
      <ul>
         <li><router-link to="/">DPG Statistics</router-link></li>
         <li><router-link to="/reports">Digitization Reports</router-link></li>
         <li><router-link to="/audit">Master File Audit</router-link></li>
      </ul>
   </nav>
   <router-view />
</template>

<script setup>
import UvaLibraryLogo from './components/UvaLibraryLogo.vue'
import {useSystemStore} from '@/stores/system'
import {useReportStore} from '@/stores/reporting'
import { onMounted } from 'vue'

const systemStore = useSystemStore()
const reportStore = useReportStore()

onMounted( () => {
   systemStore.getVersion()
   systemStore.getWorkflows()
   reportStore.init()
})
</script>


<style lang="scss">
@import "@/assets/base.scss";
header {
   background-color: var(--uvalib-brand-blue);
   color: white;
   padding: 1vw 20px;
   text-align: left;
   position: relative;
   .main-header {
      box-sizing: border-box;
      display: flex;
      flex-direction: row;
      flex-wrap: nowrap;
      justify-content: space-between;
      align-content: stretch;
      align-items: center;
   }
   div.library-link {
      width: 220px;
      order: 0;
      flex: 0 1 auto;
      align-self: flex-start;
   }

   div.site-link {
      order: 0;
      font-size: 1.5em;
      a {
         color: white;
         text-decoration: none;
         &:hover {
            text-decoration: underline;
         }
      }
   }

   p.version {
      margin: 5px 0 0 0;
      font-size: 0.5em;
      text-align: right;
   }
}
 .menu {
   background-color: var(--uvalib-blue-alt-darkest);
   ul {
      display: block;
      position: relative;
      list-style: none;
      margin: 0;
      padding: 5px;
      text-align: left;
      li {
         display: inline-block;
         padding: 5px 10px;
         margin: 0;
         font-weight: 500;
         position: relative;
         a {
            color: #aaa;
            font-weight: bold;
            &:hover {
               color: white;
            }
         }
         a.router-link-active {
           // color: var(--uvalib-teal-light);
           color: white;
         }
      }
   }
}
</style>
