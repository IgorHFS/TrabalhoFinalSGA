package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sga/models"
	"sga/services"
)

type AlocacaoHandler struct{ Service services.AlocacaoService }

func (h AlocacaoHandler) Create(c *gin.Context) {
	var in models.AlocacaoInput
	if !bindJSON(c, &in) {
		return
	}
	if e := h.Service.Create(c.Request.Context(), c.Param("id"), in); e != nil {
		respondError(c, e)
		return
	}
	c.JSON(http.StatusOK, gin.H{"mensagem": "turma alocada com sucesso"})
}
