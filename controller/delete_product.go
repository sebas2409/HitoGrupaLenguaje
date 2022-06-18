package controller

import (
	"HitoGrupaLenguaje/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h handler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Productos

	if result := h.DB.First(&product, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&product)

	c.JSON(http.StatusOK, gin.H{
		"Mensaje": "Producto eliminado",
	})
}
