package handlers

import (
	"net/http"

	"github.com/filipebuba/restaurante/internal/core/domain"
	"github.com/gin-gonic/gin"
)

var foodItems []domain.FoodItem

// Handlers para Food Items
func getFoodItems(c *gin.Context) {
	c.JSON(http.StatusOK, foodItems)
}

func createFoodItem(c *gin.Context) {
	var newFoodItem domain.FoodItem
	if err := c.ShouldBindJSON(&newFoodItem); err == nil {
		foodItems = append(foodItems, newFoodItem)
		c.JSON(http.StatusCreated, newFoodItem)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}
