package main

import (
	"evento-api/database" // Importa o pacote database
	"evento-api/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializa a conexão com o banco de dados
	database.Connect()

	r := gin.Default()

	// Configuração do middleware CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Liberado somente para desenvolvimento
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Configuração das rotas
	routes.SetupRoutes(r)

	// Inicializa o servidor na porta 8000
	r.Run(":8000")
}
