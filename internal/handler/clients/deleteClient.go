package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) DeleteCliente(c *gin.Context) {
	id := c.Param("id")
	err := h.clientService.DeleteCliente(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}