package controller

import (
	"HitoGrupaLenguaje/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AddProductModel struct {
	Nombre      string  `json:"nombre"`
	Precio      float64 `json:"precio"`
	Unidades    int     `json:"unidades"`
	FechaCompra string  `json:"fechaCompra"`
}

func (h handler) AddProduct(c *gin.Context) {
	body := AddProductModel{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	fmt.Println(body)
	var product models.Productos

	product.Nombre = body.Nombre
	product.Precio = body.Precio
	product.Unidades = body.Unidades
	product.FechaCompra = body.FechaCompra

	if result := h.DB.Create(&product); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &product)
}
