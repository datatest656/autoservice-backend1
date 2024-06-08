package controllers

import (
	"autoservice-backend/internal"
	"autoservice-backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateFeedback godoc
// @Summary Create feedback
// @Description Create a new feedback
// @Tags Feedback
// @Accept  json
// @Produce  json
// @Param input body models.Feedback true "Feedback info"
// @Success 200 {object} models.Feedback
// @Failure 400 {object} models.ErrorResponse
// @Router /api/feedback [post]
func CreateFeedback(c *gin.Context) {
	var input models.Feedback
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}
	internal.DB.Create(&input)
	c.JSON(http.StatusOK, models.SuccessResponse{Data: input})
}
