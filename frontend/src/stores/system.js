import { defineStore } from 'pinia'
import axios from 'axios'

export const useSystemStore = defineStore('system', {
	state: () => ({
		version: "unknown",
		workflows: [],
		error: "",
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
		getWorkflows() {
         axios.get("/api/workflows").then(response => {
            this.workflows = response.data
         }).catch(e => {
            this.error = e
         })
      },
	}
})