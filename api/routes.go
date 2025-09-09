package api

import (
	"log"

	"github.com/Hertucktor/authapi/dbhandler"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() {

	router := gin.Default()

	router.GET("/status", getStatus)
	// router.Get("/health", getHealth) for Overall health stats on the db
	// router.Get("/metrics", getMetrics) Uptime, Resource Usage etc.
	router.POST("/register", dbhandler.RegisterUser)
	router.POST("/login", dbhandler.LoginUser)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server konnte nicht gestartet werden: ", err)
	}
}
