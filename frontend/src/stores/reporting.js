import { defineStore } from 'pinia'
import axios from 'axios'

export const useReportStore = defineStore('report', {
	state: () => ({
		version: "unknown",
		error: "",
		dateRangeType: "before",
		startDate: new Date(),
		endDate: null,
		imageStats: {
			total: 0,
			DL: 0,
			DPLA: 0,
			rangeText: "",
			loading: false,
		},
		storageStats: {
			all: 0,
			DL: 0,
			rangeText: "",
			loading: false,
		},
		metadataStats: {
			total: 0,
			sirsi: 0,
			xml: 0,
			totalDL: 0,
			sirsiDL: 0,
			xmlDL: 0,
			totalDPLA: 0,
			sirsiDPLA: 0,
			xmlDPLA: 0,
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
            this.error = e
         })
      },

		getImageSats() {
			let dateParam = getDateParam(this.dateRangeType, this.startDate, this.endDate)
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
            this.error = e
				this.imageStats.loading = false
         })
		},

		getStorageSats() {
			let dateParam = getDateParam(this.dateRangeType, this.startDate, this.endDate)
			let url = "/api/stats/storage"
			if (dateParam != "") {
				url += "?date="+encodeURIComponent(dateParam)
			}
			this.storageStats.rangeText = dateParam
			this.storageStats.loading = true
			axios.get(url).then(response => {
				this.storageStats.all = response.data.all
				this.storageStats.DL = response.data.dl
				this.storageStats.loading = false
			}).catch(e => {
            this.error = e
				this.storageStats.loading = false
         })
		},

		getMetadataSats() {
			let dateParam = getDateParam(this.dateRangeType, this.startDate, this.endDate)
			let url = "/api/stats/metadata"
			if (dateParam != "") {
				url += "?date="+encodeURIComponent(dateParam)
			}
			this.metadataStats.rangeText = dateParam
			this.metadataStats.loading = true
			axios.get(url).then(response => {
				this.metadataStats.total = response.data.all.total
				this.metadataStats.sirsi = response.data.all.sirsi
				this.metadataStats.xml = response.data.all.xml
				this.metadataStats.totalDL = response.data.DL.total
				this.metadataStats.sirsiDL = response.data.DL.sirsi
				this.metadataStats.xmlDL = response.data.DL.xml
				this.metadataStats.totalDPLA = response.data.DPLA.total
				this.metadataStats.sirsiDPLA = response.data.DPLA.sirsi
				this.metadataStats.xmlDPLA = response.data.DPLA.xml
				this.metadataStats.loading = false
			}).catch(e => {
            this.error = e
				this.metadataStats.loading = false
         })
		}
	}
})

function getDateParam(rangeType, startDate, endDate) {
	let dateParam = ""
	if (rangeType == "before") {
		dateParam = `BEFORE ${dateString(startDate)}`
	} else if (rangeType == "after") {
		dateParam = `AFTER ${dateString(startDate)}`
	} else {
		dateParam = `${dateString(startDate)} TO ${dateString(endDate)}`
	}
	return dateParam
}

function dateString(date) {
	let mo = `${date.getMonth()+1}`
	let day = `${date.getDate()}`
	return `${date.getFullYear()}-${mo.padStart(2,'0')}-${day.padStart(2,'0')}`
}
