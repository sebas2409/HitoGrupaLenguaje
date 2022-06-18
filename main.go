package main

import (
	"HitoGrupaLenguaje/controller"
	"HitoGrupaLenguaje/database"
	"github.com/gin-gonic/gin"
)

func main() {

	h := database.GetConnection()
	r := gin.Default()

	controller.RegisterRoutes(r, h)

	err := r.Run()
	if err != nil {
		return
	}
}
