package handlers

import (
	"net/http"

	"github.com/filipebuba/restaurante/internal/core/domain"
	"github.com/gin-gonic/gin"
)

var tables []domain.Table

// Handlers para Tables
func getTables(c *gin.Context) {
	c.JSON(http.StatusOK, tables)
}

func createTable(c *gin.Context) {
	var newTable domain.Table
	if err := c.ShouldBindJSON(&newTable); err == nil {
		tables = append(tables, newTable)
		c.JSON(http.StatusCreated, newTable)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}
