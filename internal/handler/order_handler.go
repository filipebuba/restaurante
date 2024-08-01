package handlers

import (
	"net/http"

	"github.com/filipebuba/restaurante/internal/core/domain"
	"github.com/gin-gonic/gin"
)

var orders []domain.Order

// Handlers para Orders
func getOrders(c *gin.Context) {
	c.JSON(http.StatusOK, orders)
}

func createOrder(c *gin.Context) {
	var newOrder domain.Order
	if err := c.ShouldBindJSON(&newOrder); err == nil {
		orders = append(orders, newOrder)
		c.JSON(http.StatusCreated, newOrder)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}
