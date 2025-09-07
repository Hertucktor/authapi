package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Status struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Version   string    `json:"version"`
}

func getStatus(c *gin.Context) {
	status := Status{
		Status:    "ok",
		Timestamp: time.Now(),
		Version:   "0.0.1",
	}

	c.IndentedJSON(http.StatusOK, status)
}

/* future expansion
- get Status from DB
*/
