package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (svc *serviceContext) getImageStats(c *gin.Context) {
	// expected query format: start=2022-02-17&end=2022-02-28
	log.Printf("INFO: get image statistics")
	startDateStr := c.Query("start")
	endDateStr := c.Query("end")

	var imageResp struct {
		All  int64 `json:"all"`
		DL   int64 `json:"dl"`
		DPLA int64 `json:"dpla"`
	}

	// all, dl, dpla
	for i := 0; i < 3; i++ {
		cntQuery := svc.GDB.Debug().Table("master_files")
		if startDateStr != "" {
			if endDateStr != "" {
				cntQuery = cntQuery.Where("master_files.created_at >= ? and master_files.created_at <= ?", startDateStr, endDateStr)
				log.Printf("INFO: get image counts between %s - %s", startDateStr, endDateStr)
			} else {
				log.Printf("INFO: get image counts before %s", startDateStr)
				cntQuery = cntQuery.Where("master_files.created_at <== ?", startDateStr)
			}
		}

		count := &imageResp.All
		if i == 1 {
			count = &imageResp.DL
			cntQuery = cntQuery.Where("master_files.date_dl_ingest is not null")
			log.Printf("INFO: get DL images")
		} else if i == 2 {
			count = &imageResp.DPLA
			cntQuery = cntQuery.Joins("inner join metadata m on metadata_id=m.id and m.dpla=1")
			log.Printf("INFO: get DPLA images")
		} else {
			log.Printf("INFO: get all images")
		}

		err := cntQuery.Count(count).Error
		if err != nil {
			log.Printf("ERROR: unable to image counts: %s", err.Error())
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, imageResp)
}
