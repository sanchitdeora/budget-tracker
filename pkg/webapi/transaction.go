package webapi

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

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

func (service *ApiService) GetAllTransactionsByDate(c *gin.Context) {

	startEpoch := c.Param("startEpoch")
	endEpoch := c.Param("endEpoch")

	t1, err := strconv.ParseInt(startEpoch, 10, 64)
	t2, err := strconv.ParseInt(endEpoch, 10, 64)

	fmt.Println("Start time", time.UnixMilli(t1), "end time: ", time.UnixMilli(t2))
	// fmt.Println(time.Unix(startEpoch, 0))

	var response []models.Transaction

	response, err = service.TransactionService.GetTransactionsByDate(c, time.UnixMilli(t1), time.UnixMilli(t2))
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