/* trunk-ignore-all(gofmt) */
package handlers

import (
	"net/http"

	"github.com/filipebuba/restaurante/internal/core/domain"
	"github.com/gin-gonic/gin"
)


func (h *handler) CreateCliente(c *gin.Context) {
    var newCliente cliente

    if err := c.ShouldBindJSON(&newCliente); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON provided", "details": err.Error()})
        return
    }

    clientResult := domain.Cliente{
        Nome:      newCliente.Nome,
        Telefone:  newCliente.Telefone,
        Email:     newCliente.Email,
        Feedbacks: newCliente.Feedbacks,
        Orders:    newCliente.Orders,
    }

    client, err := h.clientService.CreateCliente(c.Request.Context(), clientResult)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create client", "details": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, client)
}






