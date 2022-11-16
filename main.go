package Avito

import (
	"Avito/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/balance/:user_id", getBalance)
	router.POST("/balance", addBalance)
	router.Run("localhost:8083")
}

func getBalance(c *gin.Context) { //функция получения баланса пользователя
	user_id := c.Param("user_id")

	balance := models.GetBalance(user_id)

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
