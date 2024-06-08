package controllers

import (
	"autoservice-backend/internal"
	"autoservice-backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetServices godoc
// @Summary Get all services
// @Description Get all services
// @Tags Services
// @Produce json
// @Success 200 {array} models.Service
// @Failure 500 {object} models.ErrorResponse
// @Router /api/services [get]
func GetServices(c *gin.Context) {
	var services []models.Service
	internal.DB.Find(&services)
	c.JSON(http.StatusOK, models.SuccessResponse{Data: services})
}

// GetServiceByID godoc
// @Summary Get service by ID
// @Description Get service by ID
// @Tags Services
// @Produce json
// @Param id path uint true "Service ID"
// @Success 200 {object} models.Service
// @Failure 404 {object} models.ErrorResponse
// @Router /api/services/{id} [get]
func GetServiceByID(c *gin.Context) {
	id := c.Param("id")
	var service models.Service
	if err := internal.DB.First(&service, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Service not found"})
		return
	}
	c.JSON(http.StatusOK, models.SuccessResponse{Data: service})
}

// CreateService godoc
// @Summary Create a new service
// @Description Create a new service
// @Tags Services
// @Accept  json
// @Produce  json
// @Param input body models.Service true "Service info"
// @Success 200 {object} models.Service
// @Failure 400 {object} models.ErrorResponse
// @Router /api/admin/services [post]
func CreateService(c *gin.Context) {
	var input models.Service
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}
	internal.DB.Create(&input)
	c.JSON(http.StatusOK, models.SuccessResponse{Data: input})
}

// UpdateService godoc
// @Summary Update a service
// @Description Update a service by ID
// @Tags Services
// @Accept  json
// @Produce  json
// @Param id path uint true "Service ID"
// @Param input body models.Service true "Service info"
// @Success 200 {object} models.Service
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /api/admin/services/{id} [put]
func UpdateService(c *gin.Context) {
	id := c.Param("id")
	var service models.Service
	if err := internal.DB.First(&service, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Service not found"})
		return
	}

	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	internal.DB.Save(&service)
	c.JSON(http.StatusOK, models.SuccessResponse{Data: service})
}

// DeleteService godoc
// @Summary Delete a service
// @Description Delete a service by ID
// @Tags Services
// @Accept  json
// @Produce  json
// @Param id path uint true "Service ID"
// @Success 200 {object} models.SuccessResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /api/admin/services/{id} [delete]
func DeleteService(c *gin.Context) {
	id := c.Param("id")
	var service models.Service
	if err := internal.DB.First(&service, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Service not found"})
		return
	}

	internal.DB.Delete(&service)
	c.JSON(http.StatusOK, models.SuccessResponse{Data: true})
}
