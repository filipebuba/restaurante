package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetClienteByID(c *gin.Context) {
	clienteID := c.Param("id")

	cliente, err := h.clientService.GetClienteByID(c, clienteID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if cliente == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cliente not found"})
		return
	}

	c.JSON(http.StatusOK, cliente)
}