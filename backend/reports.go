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
	err := svc.GDB.Joins("inner join units u on order_id = orders.id").
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

func (svc *serviceContext) getProductivityReport(c *gin.Context) {
	log.Printf("INFO: get productivity report")
	workflowID := c.Query("workflow")
	startDate := c.Query("start")
	endDate := c.Query("end")

	var resp struct {
		CompletedProjects int      `json:"completedProjects"`
		Types             []string `json:"types"`
		Productivity      []int64  `json:"productivity"`
	}
	type productivityRec struct {
		Type  string
		Count int64
	}
	var dbData []productivityRec
	err := svc.GDB.Table("projects").Select("c.name as type, count(projects.id) as count").
		Joins("inner join categories c on c.id = category_id").
		Where("workflow_id=?", workflowID).
		Where("finished_at >= ? and finished_at <= ?", startDate, endDate).
		Group("c.id").Find(&dbData).Error
	if err != nil {
		log.Printf("ERROR: unable to get productivity report: %s", err.Error())
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	for _, p := range dbData {
		resp.CompletedProjects += int(p.Count)
		resp.Types = append(resp.Types, p.Type)
		resp.Productivity = append(resp.Productivity, p.Count)
	}

	c.JSON(http.StatusOK, resp)
}

func (svc *serviceContext) getProblemsReport(c *gin.Context) {
	log.Printf("INFO: get problems report")
	workflowID := c.Query("workflow")
	startDate := c.Query("start")
	endDate := c.Query("end")

	var resp struct {
		Types    []string `json:"types"`
		Problems []int64  `json:"problems"`
	}
	type problemRec struct {
		Type  string
		Count int64
	}

	var dbData []problemRec
	err := svc.GDB.Table("notes").Select("pb.label as type, count(notes.id) as count").
		Joins("inner join notes_problems np on np.note_id = notes.id").
		Joins("inner join problems pb on pb.id = np.problem_id").
		Joins("inner join projects p on project_id = p.id").
		Where("note_type=?", 2).Where("pb.label <> 'Finalization'").
		Where("finished_at is not null").Where("workflow_id=?", workflowID).
		Where("finished_at >= ? and finished_at <= ?", startDate, endDate).
		Group("problem_id").Find(&dbData).Error
	if err != nil {
		log.Printf("ERROR: unable to get productivity report: %s", err.Error())
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	for _, p := range dbData {
		resp.Types = append(resp.Types, p.Type)
		resp.Problems = append(resp.Problems, p.Count)
	}

	c.JSON(http.StatusOK, resp)
}

func (svc *serviceContext) getPageTimesReport(c *gin.Context) {
	log.Printf("INFO: get average page times report")
	workflowID := c.Query("workflow")
	startDate := c.Query("start")
	endDate := c.Query("end")

	log.Printf("INFO: get all catagories")
	type category struct {
		ID   int64
		Name string
	}
	var categories []category
	err := svc.GDB.Where("name not like ?", "Atiz%").Find(&categories).Error
	if err != nil {
		log.Printf("ERROR: unable to get categories: " + err.Error())
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	log.Printf("INFO: get unit timing")
	type timingRec struct {
		ProjectID int64
		UnitID    int64
		Category  string
		TotalMins int64
	}

	var timings []timingRec
	err = svc.GDB.Table("projects").Select("projects.id as project_id, projects.unit_id as unit_id, c.name as category, sum(duration_minutes) as total_mins").
		Joins("inner join assignments a on projects.id = a.project_id").
		Joins("inner join categories c on c.id = projects.category_id").
		Where("workflow_id=?", workflowID).
		Where("projects.finished_at >= ?", startDate).
		Where("projects.finished_at <= ?", endDate).
		Group("projects.id").Find(&timings).Error
	if err != nil {
		log.Printf("ERROR: unable to get project timing report: %s", err.Error())
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	log.Printf("INFO: unit master file counts")
	type unitImageCountRec struct {
		ProjectID int64
		UnitID    int64
		Images    int64
	}

	var unitImageCounts []unitImageCountRec
	err = svc.GDB.Table("projects").Select("projects.id as project_id, projects.unit_id as unit_id, count(f.id) as images").
		Joins("inner join units u on projects.unit_id = u.id").
		Joins("inner join master_files f on f.unit_id = u.id").
		Where("workflow_id=?", workflowID).
		Where("projects.finished_at >= ?", startDate).
		Where("projects.finished_at <= ?", endDate).
		Group("u.id").Find(&unitImageCounts).Error
	if err != nil {
		log.Printf("ERROR: unable to get unit image count report: %s", err.Error())
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	log.Printf("INFO: generate report from collected stats")
	type Stats struct {
		TotalMins   int64   `json:"mins"`
		TotalImages int64   `json:"images"`
		TotalUnits  int64   `json:"units"`
		AvgPageTime float64 `json:"avgPageTime"`
	}

	// init response with blank stats rec for each category
	resp := make(map[string]*Stats)
	for _, c := range categories {
		resp[c.Name] = &Stats{}
	}

	// sum up unit / image counts for each category
	for _, t := range timings {
		tgtStats := resp[t.Category]
		for _, u := range unitImageCounts {
			if u.UnitID == t.UnitID {
				tgtStats.TotalUnits++
				tgtStats.TotalImages += u.Images
				tgtStats.TotalMins += t.TotalMins
				break
			}
		}
	}

	// calculate average page time for each category
	for _, stats := range resp {
		if stats.TotalImages > 0 {
			stats.AvgPageTime = float64(stats.TotalMins) / float64(stats.TotalImages)
		}
	}

	c.JSON(http.StatusOK, resp)
}
