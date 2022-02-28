package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (svc *serviceContext) getImageStats(c *gin.Context) {
	log.Printf("INFO: get image statistics")
	dateQStr := c.Query("date")
	if (strings.Contains(dateQStr, "TO") || strings.Contains(dateQStr, "AFTER") || strings.Contains(dateQStr, "BEFORE")) == false {
		log.Printf("ERROR: invalid date query [%s]", dateQStr)
		c.String(http.StatusBadRequest, fmt.Sprintf("%s is not valid", dateQStr))
		return
	}

	var imageResp struct {
		All  int64 `json:"all"`
		DL   int64 `json:"dl"`
		DPLA int64 `json:"dpla"`
	}

	// all, dl, dpla
	for i := 0; i < 3; i++ {
		cntQuery := svc.GDB.Table("master_files")

		if strings.Contains(dateQStr, "TO") {
			bits := strings.Split(dateQStr, " ")
			cntQuery = cntQuery.Where("master_files.created_at >= ? and master_files.created_at <= ?", bits[0], bits[2])
		} else if strings.Contains(dateQStr, "AFTER") {
			bits := strings.Split(dateQStr, " ")
			cntQuery = cntQuery.Where("master_files.created_at >= ?", bits[1])
		} else if strings.Contains(dateQStr, "BEFORE") {
			bits := strings.Split(dateQStr, " ")
			cntQuery = cntQuery.Where("master_files.created_at <= ?", bits[1])
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

func (svc *serviceContext) getStorageStats(c *gin.Context) {
	log.Printf("INFO: get storage statistics")
	dateQStr := c.Query("date")
	if (strings.Contains(dateQStr, "TO") || strings.Contains(dateQStr, "AFTER") || strings.Contains(dateQStr, "BEFORE")) == false {
		log.Printf("ERROR: invalid date query [%s]", dateQStr)
		c.String(http.StatusBadRequest, fmt.Sprintf("%s is not valid", dateQStr))
		return
	}

	var storageResp struct {
		All float64 `json:"all"`
		DL  float64 `json:"dl"`
	}

	// all, dl
	for i := 0; i < 2; i++ {
		szQuery := svc.GDB.Debug().Table("master_files")

		if strings.Contains(dateQStr, "TO") {
			bits := strings.Split(dateQStr, " ")
			szQuery = szQuery.Where("master_files.created_at >= ? and master_files.created_at <= ?", bits[0], bits[2])
		} else if strings.Contains(dateQStr, "AFTER") {
			bits := strings.Split(dateQStr, " ")
			szQuery = szQuery.Where("master_files.created_at >= ?", bits[1])
		} else if strings.Contains(dateQStr, "BEFORE") {
			bits := strings.Split(dateQStr, " ")
			szQuery = szQuery.Where("master_files.created_at <= ?", bits[1])
		}

		sizeGB := &storageResp.All
		if i == 1 {
			sizeGB = &storageResp.DL
			szQuery = szQuery.Where("master_files.date_dl_ingest is not null")
			log.Printf("INFO: get DL images size")
		} else {
			log.Printf("INFO: get all images size")
		}

		err := szQuery.Select("sum(filesize)/1073741824.0 as size_gb").Row().Scan(sizeGB)
		if err != nil {
			log.Printf("ERROR: unable to image size: %s", err.Error())
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, storageResp)
}
