package webapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanchitdeora/budget-tracker/models"
)


func (service *ApiService) GetAllBudgets(c *gin.Context) {

	budgets, err := service.BudgetService.GetBudgets(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(200, gin.H{
		"message": "Success",
		"body":    budgets,
	})

}

func (service *ApiService) GetBudgetById(c *gin.Context) {
	
	budget, err := service.BudgetService.GetBudgetById(c, c.Param("id"))
	if budget.BudgetId == "" {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"body":    budget,
	})

}

func (service *ApiService) CreateBudgetByUser(c *gin.Context) {
	var budget models.Budget
	err := c.BindJSON(&budget)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	budgetId, err := service.BudgetService.CreateBudgetByUser(c, &budget)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(201, gin.H{
		"message": "Success",
		"body":    budgetId,
	})

}

func (service *ApiService) UpdateBudgetById(c *gin.Context) {
	
	var budget models.Budget
	budgetId := c.Param("id")
	err := c.BindJSON(&budget)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	
	budgetId, err = service.BudgetService.UpdateBudgetById(c, budgetId, &budget)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(200, gin.H{
		"message": "Success",
		"body":    budgetId,
	})

}

func (service *ApiService) DeleteBudgetById(c *gin.Context) {
	budgetId, err := service.BudgetService.DeleteBudgetById(c, c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(200, gin.H{
		"message": "Success",
		"body":    budgetId,
	})

}