package budget

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanchitdeora/budget-tracker/src/models"
)


func GetAllBudgets(c *gin.Context) {

	var response []models.Budget
	err := getBudgets(c, &response)
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
	
	var response models.Budget
	err := getBudgetById(c, c.Param("id"), &response)
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
	
	var budget models.Budget
	err := c.BindJSON(&budget)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	budgetId, err := createBudget(c, budget)
	if err != nil {
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
	
	budgetId, err := updateBudgetById(c, c.Param("id"), budget)
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
	budgetId, err := deleteBudgetById(c, c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(200, gin.H{
		"message": "Success",
		"body":    budgetId,
	})

}
