package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() {

	router := gin.Default()

	router.GET("/user/:name", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, wir sind live!")
	})

	// By default, it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
	// router.Run(":3000") for a hard coded port
}
