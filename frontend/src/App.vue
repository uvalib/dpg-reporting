<template>
   <header>
      <div class="library-link">
         <a target="_blank" href="https://library.virginia.edu">
            <UvaLibraryLogo />
         </a>
      </div>
      <div class="site-link">
         <router-link to="/">DPG Reporting</router-link>
         <p class="version">v{{ reportStore.version }}</p>
      </div>
   </header>

   <router-view />
</template>

<script>
import UvaLibraryLogo from '@/components/UvaLibraryLogo.vue'
import {useReportStore} from '@/stores/reporting'
export default {
   setup() {
      const reportStore = useReportStore()
      return {
         reportStore,
      }
   },
   components: {
      UvaLibraryLogo
   },
   created() {
      this.reportStore.getVersion()
   }
}
</script>

<style lang="scss">
@import "@/assets/base.scss";
header {
   background-color: var(--uvalib-brand-blue);
   color: white;
   padding: 1vw 20px;
   text-align: left;
   position: relative;
   box-sizing: border-box;
   display: flex;
   flex-direction: row;
   flex-wrap: nowrap;
   justify-content: space-between;
   align-content: stretch;
   align-items: center;

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
</style>
