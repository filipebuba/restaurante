package handlers

import (
	"net/http"

	"github.com/filipebuba/restaurante/internal/core/domain"
	"github.com/gin-gonic/gin"
)

var invoices []domain.Invoice

// Handlers para Invoices
func getInvoices(c *gin.Context) {
	c.JSON(http.StatusOK, invoices)
}

func createInvoice(c *gin.Context) {
	var newInvoice domain.Invoice
	if err := c.ShouldBindJSON(&newInvoice); err == nil {
		invoices = append(invoices, newInvoice)
		c.JSON(http.StatusCreated, newInvoice)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}
