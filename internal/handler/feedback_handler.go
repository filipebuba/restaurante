package handlers

import (
	"net/http"

	"github.com/filipebuba/restaurante/internal/core/domain"
	"github.com/gin-gonic/gin"
)

var feedbacks []domain.Feedback

// Handlers para Feedbacks
func getFeedbacks(c *gin.Context) {
	c.JSON(http.StatusOK, feedbacks)
}

func createFeedback(c *gin.Context) {
	var newFeedback domain.Feedback
	if err := c.ShouldBindJSON(&newFeedback); err == nil {
		feedbacks = append(feedbacks, newFeedback)
		c.JSON(http.StatusCreated, newFeedback)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}
