package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handlers para Clientes
func (h *handler) GetClientes(c *gin.Context) {
	result, err := h.clientService.GetAllClientes(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, result)
}
