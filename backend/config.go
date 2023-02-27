package main

import (
	"flag"
	"log"
)

type dbConfig struct {
	Host string
	Port int
	User string
	Pass string
	Name string
}

type configData struct {
	port     int
	db       dbConfig
	adminURL string
	iiifURL  string
}

func getConfiguration() *configData {
	var config configData
	flag.IntVar(&config.port, "port", 8080, "Port to offer service on (default 8085)")

	// DB connection params
	flag.StringVar(&config.db.Host, "dbhost", "", "Database host")
	flag.IntVar(&config.db.Port, "dbport", 3306, "Database port")
	flag.StringVar(&config.db.Name, "dbname", "", "Database name")
	flag.StringVar(&config.db.User, "dbuser", "", "Database user")
	flag.StringVar(&config.db.Pass, "dbpass", "", "Database password")
	flag.StringVar(&config.adminURL, "tsadmin", "https://tracksys.lib.virginia.edu/admin", "URL for TrackSys ADMIN interface")
	flag.StringVar(&config.iiifURL, "iiif", "https://iiif.lib.virginia.edu/iiif", "URL UVA IIIF server")

	flag.Parse()

	if config.db.Host == "" {
		log.Fatal("Parameter dbhost is required")
	}
	if config.db.Name == "" {
		log.Fatal("Parameter dbname is required")
	}
	if config.db.User == "" {
		log.Fatal("Parameter dbuser is required")
	}
	if config.db.Pass == "" {
		log.Fatal("Parameter dbpass is required")
	}

	log.Printf("[CONFIG] port          = [%d]", config.port)
	log.Printf("[CONFIG] iiif          = [%s]", config.iiifURL)
	log.Printf("[CONFIG] tsadmi        = [%s]", config.adminURL)
	log.Printf("[CONFIG] port          = [%d]", config.port)
	log.Printf("[CONFIG] dbhost        = [%s]", config.db.Host)
	log.Printf("[CONFIG] dbport        = [%d]", config.db.Port)
	log.Printf("[CONFIG] dbname        = [%s]", config.db.Name)
	log.Printf("[CONFIG] dbuser        = [%s]", config.db.User)

	return &config
}
