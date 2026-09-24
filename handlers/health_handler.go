package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type HealthHandler struct{ Version string }

func (h HealthHandler) Check(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "healthy", "timestamp": time.Now().UTC(), "version": h.Version})
}
