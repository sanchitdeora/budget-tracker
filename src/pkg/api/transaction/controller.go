package transaction

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanchitdeora/budget-tracker/src/models"
)

func GetAllTransactions(c *gin.Context) {

	var response []models.Transaction
	err := GetTransactions(c, &response)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(200, gin.H{
		"message": "Success",
		"body":    response,
	})

}

func GetSingleTransaction(c *gin.Context) {

	var response models.Transaction
	err := GetTransactionById(c, c.Param("id"), &response)
	if response.TransactionId == "" {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(200, gin.H{
		"message": "Success",
		"body":    response,
	})

}

func CreateTransaction(c *gin.Context) {

	var transaction models.Transaction
	err := c.BindJSON(&transaction)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	transactionId, err := CreateTransactionRecord(c, transaction)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
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
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	
	transactionId, err := UpdateTransactionById(c, c.Param("id"), transaction)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	
	c.JSON(200, gin.H{
		"message": "Success",
		"body":    transactionId,
	})

}

func DeleteTransaction(c *gin.Context) {

	transactionId, err := DeleteTransactionById(c, c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(200, gin.H{
		"message": "Success",
		"body":    transactionId,
	})

}

