package handlers

import (
	"net/http"

	"github.com/filipebuba/restaurante/internal/core/domain"
	"github.com/gin-gonic/gin"
)

var promocoes []domain.Promocao

// Handlers para Promoções
func getPromocoes(c *gin.Context) {
	c.JSON(http.StatusOK, promocoes)
}

func createPromocao(c *gin.Context) {
	var newPromocao domain.Promocao
	if err := c.ShouldBindJSON(&newPromocao); err == nil {
		promocoes = append(promocoes, newPromocao)
		c.JSON(http.StatusCreated, newPromocao)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}
