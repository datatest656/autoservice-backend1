package controllers

import (
	"autoservice-backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetStatistics godoc
// @Summary Get basic statistics
// @Description Get basic statistics for the auto service
// @Tags Statistics
// @Produce json
// @Success 200 {object} models.SuccessResponse{data=models.Statistics}
// @Failure 500 {object} models.ErrorResponse
// @Router /api/admin/statistics [get]
func GetStatistics(c *gin.Context) {
	// Пример: базовая статистика
	statistics := models.Statistics{
		TotalVisits:       1000,
		TotalAppointments: 50,
	}

	c.JSON(http.StatusOK, models.SuccessResponse{Data: statistics})
}
