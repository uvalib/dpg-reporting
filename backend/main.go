package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// Version of the service
const Version = "1.2.0"

func main() {
	// Load cfg
	log.Printf("===> DPG Reporting Service is starting up <===")
	cfg := getConfiguration()
	svc := initializeService(Version, cfg)

	// Set routes and start server
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	router := gin.Default()

	// Set routes and start server
	router.Use(cors.Default())
	router.GET("/version", svc.getVersion)
	router.GET("/healthcheck", svc.healthCheck)
	api := router.Group("/api")
	{
		api.GET("/stats/archive", svc.getArchiveStats)
		api.GET("/stats/images", svc.getImageStats)
		api.GET("/stats/metadata", svc.getMetadataStats)
		api.GET("/stats/published", svc.getPublishedStats)
		api.GET("/stats/storage", svc.getStorageStats)

		api.GET("workflows", svc.getWorkflows)
		api.GET("/reports/deliveries", svc.getDeliveriesReport)
		api.GET("/reports/productivity", svc.getProductivityReport)
		api.GET("/reports/problems", svc.getProblemsReport)
		api.GET("/reports/pagetimes", svc.getPageTimesReport)
		api.GET("/reports/rejections", svc.getRejectionsReport)
		api.GET("/reports/rates", svc.getRatesReport)
	}

	// Note: in dev mode, this is never actually used. The front end is served
	// by yarn and it proxies all requests to the API to the routes above
	router.Use(static.Serve("/", static.LocalFile("./public", true)))

	// add a catchall route that renders the index page.
	// based on no-history config setup info here:
	//    https://router.vuejs.org/guide/essentials/history-mode.html#example-server-configurations
	router.NoRoute(func(c *gin.Context) {
		c.File("./public/index.html")
	})

	portStr := fmt.Sprintf(":%d", cfg.port)
	log.Printf("INFO: start DPG Reporting Service on port %s with CORS support enabled", portStr)
	log.Fatal(router.Run(portStr))
}
