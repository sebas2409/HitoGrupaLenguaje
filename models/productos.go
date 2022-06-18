package models

type Productos struct {
	Id          int     `json:"id" gorm:"primary_key;auto_increment"`
	Nombre      string  `json:"nombre"`
	Precio      float64 `json:"precio"`
	Unidades    int     `json:"unidades"`
	FechaCompra string  `json:"fechaCompra"`
}
