package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ramnakm/account-service/database"
	"github.com/ramnakm/account-service/models"
)

func CreateAccount(c *gin.Context) {
	var account models.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account.Status = "active"
	database.DB.Create(&account)
	c.JSON(http.StatusCreated, account)
}

func GetAccount(c *gin.Context) {
	var account models.Account
	if err := database.DB.First(&account, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, account)
}
