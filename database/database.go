package database

import (
	"evento-api/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(sqlite.Open("eventos.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar com o banco de dados:", err)
	}
	log.Println("Conexão com o banco realizada com sucesso!") // <-- Adicione este log
	DB.AutoMigrate(&models.Evento{})
}
