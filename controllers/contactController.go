package controllers

import (
	"autoservice-backend/internal"
	"autoservice-backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetContactInfo godoc
// @Summary Get contact information
// @Description Get contact information of the auto service
// @Tags Contact
// @Produce json
// @Success 200 {object} models.Contact
// @Failure 404 {object} models.ErrorResponse
// @Router /api/contact [get]
func GetContactInfo(c *gin.Context) {
	var contact models.Contact
	if err := internal.DB.First(&contact).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Contact information not found"})
		return
	}
	c.JSON(http.StatusOK, models.SuccessResponse{Data: contact})
}
