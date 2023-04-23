package webapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanchitdeora/budget-tracker/models"
)


func (service *ApiService) GetAllBills(c *gin.Context) {

	var response []models.Bill
	err := service.BillService.GetBills(c, &response)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(200, gin.H{
		"message": "Success",
		"body":    response,
	})

}

func (service *ApiService) GetBillById(c *gin.Context) {

	var response models.Bill
	err := service.BillService.GetBillById(c, c.Param("id"), &response)
	if response.BillId == "" {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"body":    response,
	})

}

func (service *ApiService) CreateBill(c *gin.Context) {

	var bill models.Bill
	err := c.BindJSON(&bill)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	billId, err := service.BillService.CreateBillByUser(c, bill)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(201, gin.H{
		"message": "Success",
		"body":    billId,
	})

}

func (service *ApiService) UpdateBillById(c *gin.Context) {

	var bill models.Bill
	err := c.BindJSON(&bill)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	
	billId, err := service.BillService.UpdateBillById(c, c.Param("id"), bill)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	
	c.JSON(200, gin.H{
		"message": "Success",
		"body":    billId,
	})

}

func (service *ApiService) UpdateBillIsPaid(c *gin.Context) {

	billId, err := service.BillService.UpdateBillIsPaid(c, c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	
	c.JSON(200, gin.H{
		"message": "Success",
		"body":    billId,
	})

}

func (service *ApiService) UpdateBillIsUnpaid(c *gin.Context) {

	billId, err := service.BillService.UpdateBillIsUnpaid(c, c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	
	c.JSON(200, gin.H{
		"message": "Success",
		"body":    billId,
	})

}

func (service *ApiService) DeleteBillById(c *gin.Context) {

	billId, err := service.BillService.DeleteBillById(c, c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(200, gin.H{
		"message": "Success",
		"body":    billId,
	})

}

