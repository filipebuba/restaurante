package handlers

import (
	"net/http"

	"github.com/filipebuba/restaurante/internal/core/domain"
	"github.com/gin-gonic/gin"
)

func (h *handler) UpdateCliente(c *gin.Context) {
	id := c.Param("id")

	var cliente domain.Cliente
	if err := c.ShouldBindJSON(&cliente); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON provided", "details": err.Error()})
		return
	}

	// Adiciona o ID ao cliente para garantir que estamos atualizando o cliente correto
	cliente.ID = id

	updatedCliente, err := h.clientService.UpdateCliente(c.Request.Context(), cliente)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update client", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Client updated successfully", "client": updatedCliente})
}