package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func sum(context *gin.Context) {

	n1 := context.Query("n1")
	n2 := context.Query("n1")

	num1, err := strconv.Atoi(n1)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid first number",
		})
		return
	}

	num2, err := strconv.Atoi(n2)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid second number",
		})
		return
	}

	sum := num1 + num2

	context.JSON(http.StatusOK, gin.H{
		"the result is": sum,
	})
}

func main() {
	router := gin.Default()
	router.GET("/sum", sum)
	router.Run("localhost:3000")
}
