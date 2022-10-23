package bill

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sanchitdeora/budget-tracker/src/models"
)

func GetAllBills(c *gin.Context) {

	var response []models.Bill
	err := getBills(c, &response)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{
		"message": "Success",
		"body":    response,
	})

}

func GetBillById(c *gin.Context) {

	var response models.Bill
	err := getBillById(c, c.Param("id"), &response)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{
		"message": "Success",
		"body":    response,
	})

}

func CreateBill(c *gin.Context) {

	var bill models.Bill
	err := c.BindJSON(&bill)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error",
			"body":    err,
		})
	}
	billId, err := createBill(c, bill)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(201, gin.H{
		"message": "Success",
		"body":    billId,
	})

}

func UpdateBill(c *gin.Context) {

	var bill models.Bill
	err := c.BindJSON(&bill)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error",
			"body":    err,
		})
	}
	
	billId, err := updateBillById(c, c.Param("id"), bill)
	if err != nil {
		log.Fatal(err)
	}
	
	c.JSON(200, gin.H{
		"message": "Success",
		"body":    billId,
	})

}

func DeleteBill(c *gin.Context) {

	billId, err := deleteBillById(c, c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{
		"message": "Success",
		"body":    billId,
	})

}

