package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"sga/services"
)

func respondError(c *gin.Context, err error) {
	var e *services.Error
	if errors.As(err, &e) {
		c.JSON(e.Status, gin.H{"erro": e.Message})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"erro": "erro interno do servidor"})
}
func bindJSON(c *gin.Context, target any) bool {
	if err := c.ShouldBindJSON(target); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "corpo da requisição inválido"})
		return false
	}
	return true
}
