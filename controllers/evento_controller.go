package controllers

import (
	"evento-api/database"
	"evento-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEventos(c *gin.Context) {
	var eventos []models.Evento
	database.DB.Find(&eventos)
	c.JSON(http.StatusOK, eventos)
}

func CreateEvento(c *gin.Context) {
	var evento models.Evento
	if err := c.ShouldBindJSON(&evento); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&evento)
	c.JSON(http.StatusCreated, evento)
}

func GetEventoByID(c *gin.Context) {
	var evento models.Evento
	id := c.Param("id")
	if err := database.DB.First(&evento, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Evento não encontrado"})
		return
	}
	c.JSON(http.StatusOK, evento)
}

func DeleteEvento(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&models.Evento{}, id)
	c.Status(http.StatusNoContent)
}
