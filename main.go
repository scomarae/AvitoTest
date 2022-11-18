package main

import (
	"github.com/gin-gonic/gin"
	"github.com/scomarae/AvitoTest/models"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/balance/:user_id", getBalance)
	router.POST("/balance", accrualMoneyToBalance)
	router.POST("/reserve", reserveBalance)
	router.POST("/confirm", confirmTransaction)
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

func accrualMoneyToBalance(c *gin.Context) { //функция начисления баланса
	var accrual models.AccrualMoney

	if err := c.BindJSON(&accrual); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		models.AccrualMoneyToBalance(accrual)
		c.IndentedJSON(http.StatusCreated, accrual)
	}
}

func reserveBalance(c *gin.Context) { //функция резервирования баланса
	var rbalance models.Reserve
	if err := c.BindJSON(&rbalance); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		models.ReserveBalance(rbalance)
		c.IndentedJSON(http.StatusOK, rbalance)
	}
}

func confirmTransaction(c *gin.Context) { //функция признания выручки
	var confirm models.Reserve

	if err := c.BindJSON(&confirm); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		models.ConfirmTransaction(confirm)
		c.IndentedJSON(http.StatusOK, confirm)
	}
}
