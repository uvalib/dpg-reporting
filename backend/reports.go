package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (svc *serviceContext) getDeliveriesReport(c *gin.Context) {
	log.Printf("INFO: get deliveries report")
	tgtYear := c.Query("year")
	if tgtYear == "" {
		log.Printf("ERROR: year is required")
		c.String(http.StatusBadRequest, "year is required")
		return
	}

	// NOTE: in the response, the totals are an array of 12 counts. each count corresponds to a month
	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	var resp struct {
		Months []string `json:"months"`
		Total  []int    `json:"total"`
		Late   []int    `json:"late"`
		OnTime []int    `json:"onTime"`
	}
	resp.Months = months

	type order struct {
		ID            int64
		DateDue       time.Time
		DateCompleted time.Time
	}
	var completedOrders []order
	err := svc.GDB.Debug().Joins("inner join units u on order_id = orders.id").
		Where("intended_use_id != ?", 110).
		Where("order_status=?", "completed").
		Where("date_completed like ?", fmt.Sprintf("%s%%", tgtYear)).
		Distinct("orders.id", "date_due", "date_completed").
		Order("date_completed asc").
		Find(&completedOrders).Error
	if err != nil {
		log.Printf("ERROR: unable to get deliveries report: %s", err.Error())
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	// orders are sorted by completion date asc. track the current month and add the
	// count to the correct slot in the total/late/ontime arrays.
	// NOTE: the logic below does not account for skipped months as this should never happen
	currMonth := 0
	total := 0
	late := 0
	ontime := 0
	for _, o := range completedOrders {
		orderMonth := int(o.DateCompleted.Month())
		if currMonth != orderMonth {
			if currMonth > 0 {
				resp.Total = append(resp.Total, total)
				resp.Late = append(resp.Late, late)
				resp.OnTime = append(resp.OnTime, ontime)
			}
			total = 0
			late = 0
			ontime = 0
			currMonth = orderMonth
		}
		total++
		if o.DateCompleted.Before(o.DateDue) {
			ontime++
		} else {
			late++
		}
	}

	if total > 0 {
		resp.Total = append(resp.Total, total)
		resp.Late = append(resp.Late, late)
		resp.OnTime = append(resp.OnTime, ontime)
	}

	c.JSON(http.StatusOK, resp)
}
