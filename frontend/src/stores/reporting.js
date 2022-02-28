import { defineStore } from 'pinia'
import axios from 'axios'

export const useReportStore = defineStore('report', {
	state: () => ({
		version: "unknown",
		dateRangeType: "before",
		startDate: new Date(),
		endDate: null,
		imageStats: {
			total: 0,
			DL: 0,
			DPLA: 0,
			rangeText: "",
			loading: false,
		}
	}),
	getters: {
	},
	actions: {
		getVersion() {
         axios.get("/version").then(response => {
            this.version = `v${response.data.version}-${response.data.build}`
         }).catch(e => {
            console.error(e)
         })
      },

		getImageSats() {
			let dateParam = ""
			if (this.dateRangeType == "before") {
				dateParam = `BEFORE ${dateString(this.startDate)}`
			} else if (this.dateRangeType == "after") {
				dateParam = `AFTER ${dateString(this.startDate)}`
			} else {
				dateParam = `${dateString(this.startDate)} TO ${dateString(this.endDate)}`
			}
			let url = "/api/stats/image"
			if (dateParam != "") {
				url += "?date="+encodeURIComponent(dateParam)
			}
			this.imageStats.rangeText = dateParam
			this.imageStats.loading = true
			axios.get(url).then(response => {
				this.imageStats.total = response.data.all
				this.imageStats.DL = response.data.dl
				this.imageStats.DPLA = response.data.dpla
				this.imageStats.loading = false
			}).catch(e => {
            console.error(e)
				this.imageStats.loading = false
         })
		}
	}
})

function dateString(date) {
	let mo = `${date.getMonth()+1}`
	let day = `${date.getDate()}`
	return `${date.getFullYear()}-${mo.padStart(2,'0')}-${day.padStart(2,'0')}`
}
