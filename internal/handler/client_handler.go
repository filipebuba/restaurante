/* trunk-ignore-all(gofmt) */
package handlers

import (
	"net/http"

	"github.com/filipebuba/restaurante/internal/core/domain"
	"github.com/filipebuba/restaurante/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type handler struct {
	clientService ports.ClienteService
}

var clientes []domain.Cliente

func NewHandler(clientService ports.ClienteService) *handler {
	return &handler{
		clientService: clientService,
	}
}

// Handlers para Clientes
func (h *handler) GetClientes(c *gin.Context) {
	result, err := h.clientService.GetAllClientes(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, result)
}

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

func (h *handler) DeleteCliente(c *gin.Context) {
    id := c.Param("id")
    err := h.clientService.DeleteCliente(c.Request.Context(), id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.Status(http.StatusNoContent)
}

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
