import { defineStore } from 'pinia'
import axios from 'axios'
import {useSystemStore} from '@/stores/system'

export const useReportStore = defineStore('report', {
	state: () => ({
		systemStore: useSystemStore(),
		workflowID: 1,
		startDate: null,
		endDate: null,
		deliveries: {
			labels: [],
			datasets: [],
			loading: false,
		},
		productivity: {
			loading: false,
			labels: [],
			datasets: [],
		}
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
		getProductivityReport( workflowID, start, end ) {
			let url = `/api/reports/productivity?workflow=${workflowID}&start=${dateString(start)}&end=${dateString(end)}`
			this.productivity.loading = true
			axios.get(url).then(response => {
				this.productivity.labels = response.data.types
				let prodDataset = {data: response.data.productivity, backgroundColor: "#44aacc"}
				this.productivity.datasets.splice(0, this.deliveries.datasets.length)
				this.productivity.datasets.push(prodDataset)
				this.productivity.loading = false
			}).catch(e => {
            this.systemStore.error = e
				this.productivity.loading = false
         })
		}
	}
})

function dateString(date) {
	let mo = `${date.getMonth()+1}`
	let day = `${date.getDate()}`
	return `${date.getFullYear()}-${mo.padStart(2,'0')}-${day.padStart(2,'0')}`
}
