package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes() {

	router := gin.Default()

	router.GET("/status", getStatus)
	// router.Get("/health", getHealth) for Overall health stats on the db
	// router.Get("/metrics", getMetrics) Uptime, Resource Usage etc.

	router.Run(":8080")
}
