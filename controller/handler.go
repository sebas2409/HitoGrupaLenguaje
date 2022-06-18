package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	r.GET("/", h.GetProducts)
	r.POST("/add", h.AddProduct)
	r.POST("/:id", h.DeleteProduct)
}
