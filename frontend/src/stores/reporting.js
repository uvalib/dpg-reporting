import { defineStore } from 'pinia'
import axios from 'axios'

export const useReportStore = defineStore('report', {
	id: 'counter',
	state: () => ({
		version: "unknown",
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
	}
})
