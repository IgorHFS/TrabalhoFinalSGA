package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sga/models"
	"sga/services"
)

type SalaHandler struct{ Service services.SalaService }

func (h SalaHandler) Create(c *gin.Context) {
	var in models.Sala
	if !bindJSON(c, &in) {
		return
	}
	if err := h.Service.Create(c.Request.Context(), in); err != nil {
		respondError(c, err)
		return
	}
	c.JSON(http.StatusCreated, in)
}
func (h SalaHandler) List(c *gin.Context) {
	v, e := h.Service.List(c.Request.Context())
	if e != nil {
		respondError(c, e)
		return
	}
	c.JSON(http.StatusOK, v)
}
func (h SalaHandler) Agenda(c *gin.Context) {
	v, e := h.Service.Agenda(c.Request.Context(), c.Param("id"))
	if e != nil {
		respondError(c, e)
		return
	}
	c.JSON(http.StatusOK, v)
}
