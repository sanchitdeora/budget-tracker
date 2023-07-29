package webapi

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sanchitdeora/budget-tracker/models"
	"github.com/sanchitdeora/budget-tracker/pkg/exceptions"
)


func (service *ApiService) GetAllTransactions(c *gin.Context) {

	response, err := service.TransactionService.GetTransactions(c)
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

	response, err := service.TransactionService.GetTransactionById(c, c.Param("id"))

	if err != nil {
		if errors.Is(err, exceptions.ErrValidationError) {
			c.AbortWithError(http.StatusBadRequest, err)
		} else if errors.Is(err, exceptions.ErrTransactionNotFound) {
			c.AbortWithError(http.StatusNotFound, err)
		} else {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
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
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	t2, err := strconv.ParseInt(endEpoch, 10, 64)
	
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	log.Println("Start time: ", time.UnixMilli(t1), "end time: ", time.UnixMilli(t2))

	response, err := service.TransactionService.GetTransactionsByDate(c, time.UnixMilli(t1), time.UnixMilli(t2))
	if err != nil {
		if errors.Is(err, exceptions.ErrValidationError) {
			c.AbortWithError(http.StatusBadRequest, err)
		} else {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
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
		if errors.Is(err, exceptions.ErrValidationError) {
			c.AbortWithError(http.StatusBadRequest, err)
		} else {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
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