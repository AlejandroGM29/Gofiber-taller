package database

import (
	"fmt"
	persona "goFIber/Persona"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	// Código para abrir la conexión a la base de datos (similar al ejemplo anterior)
	dsn := "root:Sobrecarga2*@tcp(localhost:3308)/bdtodo?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("Error al conectar a la base de datos: " + err.Error())
	}

	// Tu lógica de aplicación aquí...

	db.AutoMigrate(&persona.Persona{})
	// Cerrar la conexión a la base de datos al finalizar
	fmt.Println("Database Migrated")
}

// Función para cerrar la conexión a la base de datos
func Close() {
	if sqlDB, err := db.DB(); err == nil {
		sqlDB.Close()
	}
}
