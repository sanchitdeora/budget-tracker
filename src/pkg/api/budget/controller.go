package budget

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanchitdeora/budget-tracker/src/models"
)


func GetAllBudgets(c *gin.Context) {

	var response []models.Budget
	err := GetBudgets(c, &response)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(200, gin.H{
		"message": "Success",
		"body":    response,
	})

}


func GetBudgetById(c *gin.Context) {
	
	response, err := GetBudget(c, c.Param("id"))
	if response.BudgetId == "" {
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

func CreateBudget(c *gin.Context) {
	fmt.Println("\n creating budget now")
	var budget models.Budget
	err := c.BindJSON(&budget)
	if err != nil {
		fmt.Println("\nError binding JSON", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	fmt.Println("\nprint budget", budget)
	budgetId, err := createBudgetByUser(c, budget)
	if err != nil {
		fmt.Println("\nError creating budget", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(201, gin.H{
		"message": "Success",
		"body":    budgetId,
	})

}

func UpdateBudget(c *gin.Context) {
	
	var budget models.Budget
	err := c.BindJSON(&budget)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	
	budgetId, err := UpdateBudgetById(c, c.Param("id"), budget)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	
	c.JSON(200, gin.H{
		"message": "Success",
		"body":    budgetId,
	})

}

func DeleteBudget(c *gin.Context) {
	budgetId, err := DeleteBudgetById(c, c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(200, gin.H{
		"message": "Success",
		"body":    budgetId,
	})

}
