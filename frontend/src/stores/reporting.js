import { defineStore } from 'pinia'
import axios from 'axios'
import {useSystemStore} from '@/stores/system'

export const useReportStore = defineStore('report', {
	state: () => ({
		systemStore: useSystemStore(),
		deliveries: {
			months: [],
			total: [],
			onTime: [],
			late: [],
			loading: false,
		},
	}),
	getters: {
	},
	actions: {
		getDeliveriesReport( year ) {
			let url = `/api/reports/deliveries?year=${year}`
			this.deliveries.loading = true
			axios.get(url).then(response => {
				this.deliveries.months = response.data.months
				this.deliveries.total = response.data.total
				this.deliveries.onTime = response.data.onTime
				this.deliveries.late = response.data.late
				this.deliveries.loading = false
			}).catch(e => {
            this.systemStore.error = e
				this.deliveries.loading = false
         })
		},
	}
})
