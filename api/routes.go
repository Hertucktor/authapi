package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes() {

	router := gin.Default()

	router.GET("/status", getStatus)

	router.Run(":8080")
}
