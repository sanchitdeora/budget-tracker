package webapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanchitdeora/budget-tracker/models"
)


func (service *ApiService) GetAllTransactions(c *gin.Context) {

	var response []models.Transaction
	err := service.TransactionService.GetTransactions(c, &response)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(200, gin.H{
		"message": "Success",
		"body":    response,
	})

}

func (service *ApiService) GetTransactionById(c *gin.Context) {

	var response models.Transaction
	err := service.TransactionService.GetTransactionById(c, c.Param("id"), &response)
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

func (service *ApiService) CreateTransaction(c *gin.Context) {

	var transaction models.Transaction
	err := c.BindJSON(&transaction)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	transactionId, err := service.TransactionService.CreateTransaction(c, transaction)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(201, gin.H{
		"message": "Success",
		"body":    transactionId,
	})

}

func (service *ApiService) UpdateTransactionById(c *gin.Context) {

	var transaction models.Transaction
	err := c.BindJSON(&transaction)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	
	transactionId, err := service.TransactionService.UpdateTransactionById(c, c.Param("id"), transaction)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	
	c.JSON(200, gin.H{
		"message": "Success",
		"body":    transactionId,
	})

}

func (service *ApiService) DeleteTransactionById(c *gin.Context) {

	transactionId, err := service.TransactionService.DeleteTransactionById(c, c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(200, gin.H{
		"message": "Success",
		"body":    transactionId,
	})

}