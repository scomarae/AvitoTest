package main

import (
	"github.com/gin-gonic/gin"
	"github.com/scomarae/AvitoTest/models"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/balance/:user_id", getBalance)
	router.POST("/balance", addBalance)
	router.Run("localhost:8083")
}

func getBalance(c *gin.Context) { //функция получения баланса пользователя
	userId := c.Param("user_id")

	balance := models.GetBalance(userId)

	if balance == nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, balance)
	}
}

func addBalance(c *gin.Context) { //функция начисления баланса
	var balance models.UserBalance

	if err := c.BindJSON(&balance); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		models.AddBalance(balance)
		c.IndentedJSON(http.StatusCreated, balance)
	}
}
