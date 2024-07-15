 /*global process */

import { fileURLToPath, URL } from 'url'
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
   plugins: [vue()],
   resolve: {
      alias: {
         '@': fileURLToPath(new URL('./src', import.meta.url))
      }
   },
   server: { // this is used in dev mode only
      port: 8080,
      proxy: {
         '/version': {
            target: process.env.VITE_REPORTING_API, //  export VITE_REPORTING_API=http://localhost:8085
            changeOrigin: true
         },
         '/healthcheck': {
            target: process.env.VITE_REPORTING_API,
            changeOrigin: true
         },
         '/api': {
            target: process.env.VITE_REPORTING_API,
            changeOrigin: true
         },
      }
   }
})


