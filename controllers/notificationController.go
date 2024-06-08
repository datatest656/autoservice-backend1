package controllers

import (
	"autoservice-backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SendNotification godoc
// @Summary Send a notification
// @Description Send a notification via email
// @Tags Notifications
// @Accept  json
// @Produce  json
// @Param input body struct{ Email string `json:"email" binding:"required"`; Message string `json:"message" binding:"required"` } true "Notification info"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /api/notifications/send [post]
func SendNotification(c *gin.Context) {
	var input struct {
		Email   string `json:"email" binding:"required"`
		Message string `json:"message" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	// Здесь вызовите ваш сервис уведомлений для отправки SMS/Email
	// err := service.SendEmail(input.Email, input.Message)
	// if err != nil {
	//     c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send notification"})
	//     return
	// }

	// Для простоты пока отправим успешный ответ
	c.JSON(http.StatusOK, models.SuccessResponse{Data: "Notification sent successfully"})
}
