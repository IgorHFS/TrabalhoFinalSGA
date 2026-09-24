package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sga/models"
	"sga/services"
)

type TurmaHandler struct{ Service services.TurmaService }

func (h TurmaHandler) Create(c *gin.Context) {
	var in models.Turma
	if !bindJSON(c, &in) {
		return
	}
	if e := h.Service.Create(c.Request.Context(), in); e != nil {
		respondError(c, e)
		return
	}
	c.JSON(http.StatusCreated, in)
}
func (h TurmaHandler) List(c *gin.Context) {
	v, e := h.Service.List(c.Request.Context())
	if e != nil {
		respondError(c, e)
		return
	}
	c.JSON(http.StatusOK, v)
}
