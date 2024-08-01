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
	c.JSON(http.StatusOK, clientes)
}

func (h *handler) CreateCliente(c *gin.Context) {
	var newCliente domain.Cliente

	if err := c.ShouldBindJSON(&newCliente); err == nil {
		clientes = append(clientes, newCliente)
		c.JSON(http.StatusCreated, newCliente)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func (h *handler) UpdateCliente(c *gin.Context) {
	var updatedCliente domain.Cliente

	if err := c.ShouldBindJSON(&updatedCliente); err == nil {
		// Find the index of the cliente to be updated
		index := -1
		for i, cliente := range clientes {
			if cliente.ID == updatedCliente.ID {
				index = i
				break
			}
		}

		// If the cliente is found, update it
		if index != -1 {
			clientes[index] = updatedCliente
			c.JSON(http.StatusOK, updatedCliente)
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "Cliente not found"})
		}
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func (h *handler) DeleteCliente(c *gin.Context) {
	var deletedCliente domain.Cliente

	if err := c.ShouldBindJSON(&deletedCliente); err == nil {
		// Find the index of the cliente to be deleted
		index := -1
		for i, cliente := range clientes {
			if cliente.ID == deletedCliente.ID {
				index = i
				break
			}
		}

		// If the cliente is found, delete it
		if index != -1 {
			clientes = append(clientes[:index], clientes[index+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Cliente deleted successfully"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "Cliente not found"})
		}
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}
