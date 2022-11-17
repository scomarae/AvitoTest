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
	router.POST("/reserve", reserveBalance)
	//router.POST("/confirm", confirmBalance)
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

func reserveBalance(c *gin.Context) { //функция резервирования баланса
	var rbalance models.Reserve

	//нужна еще проверка на то, есть ли такой пользователь?
	if err := c.BindJSON(&rbalance); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		models.ReserveBalance(rbalance)
		c.IndentedJSON(http.StatusCreated, rbalance)
	}
}

//func confirmBalance(c *gin.Context) { //
//	var confirm models.Confirm
//
//	if err := c.BindJSON(&confirm); err != nil {
//		c.AbortWithStatus(http.StatusBadRequest)
//	} else {
//		models.ConfirmBalance(confirm)
//		c.IndentedJSON(http.StatusCreated, confirm)
//	}
//}
