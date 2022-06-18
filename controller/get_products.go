package controller

import (
	"HitoGrupaLenguaje/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h handler) GetProducts(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	var products []models.Productos
	if result := h.DB.Find(&products); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, products)
}
