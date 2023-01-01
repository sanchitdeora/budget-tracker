package goals

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanchitdeora/budget-tracker/src/models"
)


func GetAllGoals(c *gin.Context) {

	var response []models.Goal
	err := GetGoals(c, &response)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(200, gin.H{
		"message": "Success",
		"body":    response,
	})

}


func GetGoalById(c *gin.Context) {
	
	response, err := GetGoal(c, c.Param("id"))
	if response.GoalId == "" {
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

func CreateGoal(c *gin.Context) {
	
	var goal models.Goal
	err := c.BindJSON(&goal)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	goalId, err := createGoal(c, goal)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(201, gin.H{
		"message": "Success",
		"body":    goalId,
	})

}

func UpdateGoal(c *gin.Context) {
	
	var goal models.Goal
	err := c.BindJSON(&goal)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	
	goalId, err := UpdateGoalById(c, c.Param("id"), goal)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	
	c.JSON(200, gin.H{
		"message": "Success",
		"body":    goalId,
	})

}

func DeleteGoal(c *gin.Context) {
	goalId, err := DeleteGoalById(c, c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(200, gin.H{
		"message": "Success",
		"body":    goalId,
	})

}
