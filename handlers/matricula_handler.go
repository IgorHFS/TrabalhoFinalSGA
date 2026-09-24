package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sga/models"
	"sga/services"
)

type MatriculaHandler struct{ Service services.MatriculaService }

func (h MatriculaHandler) Create(c *gin.Context) {
	var in models.MatriculaInput
	if !bindJSON(c, &in) {
		return
	}
	if e := h.Service.Create(c.Request.Context(), c.Param("id"), in); e != nil {
		respondError(c, e)
		return
	}
	c.Status(http.StatusNoContent)
}
func (h MatriculaHandler) List(c *gin.Context) {
	v, e := h.Service.List(c.Request.Context(), c.Param("id"))
	if e != nil {
		respondError(c, e)
		return
	}
	c.JSON(http.StatusOK, v)
}
