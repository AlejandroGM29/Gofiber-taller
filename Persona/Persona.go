package persona

import "gorm.io/gorm"

type Persona struct {
	gorm.Model
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
	Edad   int    `json:"edad"`
	Correo string `json:"correo"`
}
