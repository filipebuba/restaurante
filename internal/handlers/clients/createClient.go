package handlers

import (
	"net/http"

	"github.com/filipebuba/restaurante/internal/core/domain"
	"github.com/gin-gonic/gin"
)


func (h *handler) CreateCliente(c *gin.Context) {
	var newCliente domain.Cliente

	if err := c.ShouldBindJSON(&newCliente); err == nil {
		//clientes = append(clientes, newCliente)
		client, err := h.clientService.CreateCliente(c, newCliente)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusCreated, client)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}






