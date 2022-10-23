package transaction

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sanchitdeora/budget-tracker/src/models"
)

func GetAllTransactions(c *gin.Context) {

	var result []models.Transaction
	result, err := getTransactions(c)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{
		"message": "Success",
		"body":    result,
	})

}

func GetTransactionById(c *gin.Context) {

	result, err := getTransactionById(c, c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{
		"message": "Success",
		"body":    result,
	})

}

func CreateTransaction(c *gin.Context) {

	var transaction models.Transaction
	err := c.BindJSON(&transaction)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error",
			"body":    err,
		})
	}
	transactionId, err := createTransaction(c, transaction)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(201, gin.H{
		"message": "Success",
		"body":    transactionId,
	})

}

func UpdateTransaction(c *gin.Context) {

	var transaction models.Transaction
	err := c.BindJSON(&transaction)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error",
			"body":    err,
		})
	}
	
	transactionId, err := updateTransactionById(c, c.Param("id"), transaction)
	if err != nil {
		log.Fatal(err)
	}
	
	c.JSON(200, gin.H{
		"message": "Success",
		"body":    transactionId,
	})

}

func DeleteTransaction(c *gin.Context) {

	transactionId, err := deleteTransactionById(c, c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{
		"message": "Success",
		"body":    transactionId,
	})

}

