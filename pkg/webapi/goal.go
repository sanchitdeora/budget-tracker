package webapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanchitdeora/budget-tracker/models"
)


func (service *ApiService) GetAllGoals(c *gin.Context) {

	goals, err := service.GoalService.GetGoals(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(200, gin.H{
		"message": "Success",
		"body":    goals,
	})

}

func (service *ApiService) GetGoalById(c *gin.Context) {
	
	goal, err := service.GoalService.GetGoalById(c, c.Param("id"))
	if goal.GoalId == "" {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"body":    goal,
	})

}

func (service *ApiService) CreateGoal(c *gin.Context) {
	
	var goal models.Goal
	err := c.BindJSON(&goal)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	goalId, err := service.GoalService.CreateGoal(c, &goal)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(201, gin.H{
		"message": "Success",
		"body":    goalId,
	})

}

func (service *ApiService) UpdateGoalById(c *gin.Context) {
	
	var goal models.Goal
	err := c.BindJSON(&goal)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	
	goalId, err := service.GoalService.UpdateGoalById(c, c.Param("id"), &goal)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	
	c.JSON(200, gin.H{
		"message": "Success",
		"body":    goalId,
	})

}

func (service *ApiService) DeleteGoalById(c *gin.Context) {
	goalId, err := service.GoalService.DeleteGoalById(c, c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(200, gin.H{
		"message": "Success",
		"body":    goalId,
	})

}