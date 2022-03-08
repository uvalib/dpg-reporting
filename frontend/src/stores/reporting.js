import { defineStore } from 'pinia'
import axios from 'axios'

export const useReportStore = defineStore('report', {
	state: () => ({
		workflowID: 1,
		startDate: null,
		endDate: null,
		deliveries: {
			labels: [],
			datasets: [],
			loading: false,
			error: ""
		},
		productivity: {
			loading: false,
			labels: [],
			datasets: [],
			totalCompleted: 0,
			error: ""
		},
		problems: {
			loading: false,
			labels: [],
			datasets: [],
			error: ""
		},
		pageTimes: {
			loading: false,
			error: "",
			labels: [],
			datasets: [],
			raw: []
		},
		rejections: {
			loading: false,
			error: "",
			data: [],
		}
	}),
	getters: {
	},
	actions: {
		init() {
			var oldDate = new Date()
			oldDate.setMonth(oldDate.getMonth() - 3)
			this.startDate  = oldDate
			this.endDate = new Date()
		},

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
				this.deliveries.error = ""
			}).catch(e => {
            this.deliveries.error = e
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
				this.productivity.totalCompleted = response.data.completedProjects
				this.productivity.loading = false
				this.productivity.error = ""
			}).catch(e => {
            this.productivity.error = e
				this.productivity.loading = false
         })
		},
		getProblemsReport( workflowID, start, end ) {
			let url = `/api/reports/problems?workflow=${workflowID}&start=${dateString(start)}&end=${dateString(end)}`
			this.problems.loading = true
			axios.get(url).then(response => {
				this.problems.labels = response.data.types
				let dataset = {data: response.data.problems, backgroundColor: "#cc4444"}
				this.problems.datasets.splice(0, this.problems.datasets.length)
				this.problems.datasets.push(dataset)
				this.problems.loading = false
				this.problems.error = ""
			}).catch(e => {
            this.problems.error = e
				this.problems.loading = false
         })
		},
		getPageTimesReport( workflowID, start, end ) {
			let url = `/api/reports/pagetimes?workflow=${workflowID}&start=${dateString(start)}&end=${dateString(end)}`
			this.pageTimes.loading = true
			axios.get(url).then(response => {
				this.pageTimes.labels.splice(0, this.pageTimes.labels.length)
				this.pageTimes.datasets.splice(0, this.pageTimes.datasets.length)
				this.pageTimes.raw.splice(0, this.pageTimes.raw.length)
				let timeDS = {data: [], backgroundColor: "#44aacc"}
				for (const [category, stats] of Object.entries(response.data)) {
					this.pageTimes.labels.push(category)
					timeDS.data.push(stats.avgPageTime)
					let row = {category: category, units: stats.units, totalMins: stats.mins,
						totalPages: stats.images, avgPageTime:  Number.parseFloat(stats.avgPageTime).toFixed(2)}
					this.pageTimes.raw.push(row)
				}
				this.pageTimes.datasets.push(timeDS)
				this.pageTimes.loading = false
				this.pageTimes.error = ""
			}).catch(e => {
            this.pageTimes.error = e
				this.pageTimes.loading = false
         })
		},
		getRejectionsReport( workflowID, start, end ) {
			let url = `/api/reports/rejections?workflow=${workflowID}&start=${dateString(start)}&end=${dateString(end)}`
			this.rejections.loading = true
			axios.get(url).then(response => {
				this.rejections.data = response.data
				this.rejections.loading = false
				this.rejections.error = ""
			}).catch(e => {
            this.rejections.error = e
				this.rejections.loading = false
         })
		},
	}
})

function dateString(date) {
	let mo = `${date.getMonth()+1}`
	let day = `${date.getDate()}`
	return `${date.getFullYear()}-${mo.padStart(2,'0')}-${day.padStart(2,'0')}`
}
