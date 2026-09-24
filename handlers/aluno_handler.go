package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sga/models"
	"sga/services"
)

type AlunoHandler struct{ Service services.AlunoService }

func (h AlunoHandler) Create(c *gin.Context) {
	var in models.Aluno
	if !bindJSON(c, &in) {
		return
	}
	if e := h.Service.Create(c.Request.Context(), in); e != nil {
		respondError(c, e)
		return
	}
	c.JSON(http.StatusCreated, in)
}
func (h AlunoHandler) List(c *gin.Context) {
	v, e := h.Service.List(c.Request.Context())
	if e != nil {
		respondError(c, e)
		return
	}
	c.JSON(http.StatusOK, v)
}
func (h AlunoHandler) Get(c *gin.Context) {
	v, e := h.Service.Get(c.Request.Context(), c.Param("id"))
	if e != nil {
		respondError(c, e)
		return
	}
	c.JSON(http.StatusOK, v)
}
