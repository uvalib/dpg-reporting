import { defineStore } from 'pinia'
import axios from 'axios'
import {useSystemStore} from '@/stores/system'

export const useReportStore = defineStore('report', {
	state: () => ({
		systemStore: useSystemStore(),
		deliveries: {
			labels: [],
			datasets: [],
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
				// convert the response data to the datastructure needed by chart.js
				this.deliveries.labels = response.data.months
				this.deliveries.datasets.splice(0, this.deliveries.datasets.length)
				var totalDataset = {data: response.data.total, backgroundColor: "#44cc44", fill: false, borderColor: "#44cc44", label: "Total", tension: 0.4}
				var okDataset = {data: response.data.onTime, backgroundColor: "#44aacc", fill: false, borderColor: "#44aacc", label: "On-Time", tension: 0.4}
				var errDataset = {data: response.data.late, backgroundColor: "#cc4444", fill: false, borderColor: "#cc4444", label: "Late", tension: 0.4}
				this.deliveries.datasets.push(totalDataset)
				this.deliveries.datasets.push(okDataset)
				this.deliveries.datasets.push(errDataset)
				this.deliveries.loading = false
			}).catch(e => {
            this.systemStore.error = e
				this.deliveries.loading = false
         })
		},
	}
})
