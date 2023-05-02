import { defineStore } from 'pinia'
import axios from 'axios'

export const useAuditStore = defineStore('audit', {
	state: () => ({
		loading: false,
		labels: [],
		datasets: [{data: [], backgroundColor: []}],
		totalAudited: 0,
		error: ""
	}),
	getters: {
	},
	actions: {
		getAuditReport( ) {
			this.loading = true
			this.labels = []
			this.datasets = [],
			this.totalAudited = 0
			axios.get(`/api/audit`).then(response => {
				let result = {data: [],
					backgroundColor: ["#44aacc","#cc4444","#cc4444","#cc4444", "#cc4444"],
				}
				response.data.results.forEach( r => {
					this.labels.push( `${r.label} (${r.total})` )
					result.data.push( r.total )
				})

				this.datasets.push(result)
				this.totalAudited = response.data.totalAudited
				this.loading = false
				this.error = ""
			}).catch(e => {
            this.error = e
				this.loading = false
         })
		},
	}
})