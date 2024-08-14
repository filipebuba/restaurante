package handlers

import (
	"net/http"

	"github.com/filipebuba/restaurante/internal/core/domain"
	"github.com/gin-gonic/gin"
)

var menus []domain.Menu

// Handlers para Menu
func getMenus(c *gin.Context) {
	c.JSON(http.StatusOK, menus)
}

func createMenu(c *gin.Context) {
	var newMenu domain.Menu
	if err := c.ShouldBindJSON(&newMenu); err == nil {
		menus = append(menus, newMenu)
		c.JSON(http.StatusCreated, newMenu)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}
