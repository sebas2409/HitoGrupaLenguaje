package database

import (
	"HitoGrupaLenguaje/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetConnection() *gorm.DB { // el nombre de las funciones van con mayusculas
	dsn := "uvxjykxrupo7xiqy:Rz4Y1ClWstrq2Bqm0P22@tcp(bgnvbnsjsyiubtvcdrcj-mysql.services.clever-cloud.com:3306)/bgnvbnsjsyiubtvcdrcj?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Imposible establecer conexión con la base de datos")
	}

	err = db.AutoMigrate(&models.Productos{})
	if err != nil {
		return nil
	}
	return db
}
