import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import DigitizationReports from '../views/DigitizationReports.vue'
import AuditReport from '../views/AuditReport.vue'

const router = createRouter({
   history: createWebHistory(import.meta.env.BASE_URL),
   routes: [
      {
         path: '/',
         name: 'home',
         component: HomeView
      },
      {
         path: '/reports',
         name: 'reports',
         component: DigitizationReports
      },
      {
         path: '/audit',
         name: 'audit',
         component: AuditReport
      },
   ]
})

export default router
