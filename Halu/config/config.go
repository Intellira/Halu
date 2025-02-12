package config

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TableSetup() (*gorm.DB, error) {
	// Database Connection
	var datasource string
	datasource = "root:root@(127.0.0.1:3306)/database_sample?parseTime=true&loc=Asia%2FJakarta"

	database, err := gorm.Open(mysql.Open(datasource), &gorm.Config{})
	if err != nil {
		log.Fatalf("Fail to Connect Database : %v", err)
	}

	sql_database, err := database.DB()
	if err != nil {
		log.Fatalf("Fail to Retrieve Database : %v", err)
	}

	sql_database.SetMaxIdleConns(10)
	sql_database.SetMaxOpenConns(100)
	sql_database.SetConnMaxLifetime(10 * time.Minute)

	log.Println("Hallo !")

	return database, nil
}
