package controllers

import (
	"autoservice-backend/internal"
	"autoservice-backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetClients godoc
// @Summary Get all clients
// @Description Get all clients
// @Tags Clients
// @Produce json
// @Success 200 {array} models.Client
// @Failure 500 {object} models.ErrorResponse
// @Router /api/clients [get]
func GetClients(c *gin.Context) {
	var clients []models.Client
	if err := internal.DB.Find(&clients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.SuccessResponse{Data: clients})
}

// GetClientByID godoc
// @Summary Get client by ID
// @Description Get client by ID
// @Tags Clients
// @Produce json
// @Param id path uint true "Client ID"
// @Success 200 {object} models.Client
// @Failure 404 {object} models.ErrorResponse
// @Router /api/clients/{id} [get]
func GetClientByID(c *gin.Context) {
	id := c.Param("id")
	var client models.Client
	if err := internal.DB.First(&client, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Client not found"})
		return
	}
	c.JSON(http.StatusOK, models.SuccessResponse{Data: client})
}

// CreateClient godoc
// @Summary Create a new client
// @Description Create a new client
// @Tags Clients
// @Accept  json
// @Produce  json
// @Param input body models.Client true "Client info"
// @Success 200 {object} models.Client
// @Failure 400 {object} models.ErrorResponse
// @Failure 409 {object} models.ErrorResponse
// @Router /api/clients [post]
func CreateClient(c *gin.Context) {
	var input models.Client
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	// Проверка на дублирование по номеру телефона или email
	var existingClient models.Client
	if err := internal.DB.Where("phone_number = ? OR email = ?", input.PhoneNumber, input.Email).First(&existingClient).Error; err == nil {
		c.JSON(http.StatusConflict, models.ErrorResponse{Error: "Client with this phone number or email already exists"})
		return
	}

	internal.DB.Create(&input)
	c.JSON(http.StatusOK, models.SuccessResponse{Data: input})
}

// UpdateClient godoc
// @Summary Update a client
// @Description Update a client by ID
// @Tags Clients
// @Accept  json
// @Produce  json
// @Param id path uint true "Client ID"
// @Param input body models.Client true "Client info"
// @Success 200 {object} models.Client
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /api/clients/{id} [put]
func UpdateClient(c *gin.Context) {
	id := c.Param("id")
	var client models.Client
	if err := internal.DB.First(&client, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Client not found"})
		return
	}

	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	internal.DB.Save(&client)
	c.JSON(http.StatusOK, models.SuccessResponse{Data: client})
}

// DeleteClient godoc
// @Summary Delete a client
// @Description Delete a client by ID
// @Tags Clients
// @Accept  json
// @Produce  json
// @Param id path uint true "Client ID"
// @Success 200 {object} models.SuccessResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /api/clients/{id} [delete]
func DeleteClient(c *gin.Context) {
	id := c.Param("id")
	var client models.Client
	if err := internal.DB.First(&client, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Client not found"})
		return
	}

	internal.DB.Delete(&client)
	c.JSON(http.StatusOK, models.SuccessResponse{Data: true})
}
