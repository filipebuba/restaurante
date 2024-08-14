package handlers

import (
	"net/http"

	"github.com/filipebuba/restaurante/internal/core/domain"
	"github.com/gin-gonic/gin"
)

var funcionarios []domain.Funcionario

// Handlers para Funcionários
func GetFuncionarios(c *gin.Context) {
	c.JSON(http.StatusOK, funcionarios)
}

func CreateFuncionario(c *gin.Context) {
	var newFuncionario domain.Funcionario
	if err := c.ShouldBindJSON(&newFuncionario); err == nil {
		funcionarios = append(funcionarios, newFuncionario)
		c.JSON(http.StatusCreated, newFuncionario)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}
