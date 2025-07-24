package main

import (
	"evento-api/database"
	"evento-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.Connect()
	routes.SetupRoutes(r)
	r.Run(":8000")
}
