package main

import (
	"log"

	"github.com/Kelniit/Halu/config"
	"github.com/Kelniit/Halu/entities"
	"github.com/Kelniit/Halu/router"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initiate Route
	r := gin.Default()
	// Test
	database, err := config.TableSetup()
	if err != nil {
		log.Fatalf("Fail to Connect Database !")
	}
	// Migrate
	errto := database.AutoMigrate(&entities.UserEntity{})
	if errto != nil {
		log.Fatalf("Migration Fail : %v", errto)
	}
	// Test Post
	router.UserRouter(r)
	// Run Server
	r.Run(":8080")
}
