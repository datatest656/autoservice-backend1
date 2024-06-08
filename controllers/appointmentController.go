package controllers

import (
	"autoservice-backend/internal"
	"autoservice-backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// CreateAppointment godoc
// @Summary Create a new appointment
// @Description Create a new appointment
// @Tags Appointments
// @Accept  json
// @Produce  json
// @Param input body models.Appointment true "Appointment info"
// @Success 200 {object} models.Appointment
// @Failure 400 {object} models.ErrorResponse
// @Router /api/appointments [post]
func CreateAppointment(c *gin.Context) {
	var input models.Appointment
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	// Если `AppointmentDate` не задано, установите текущее время
	if input.AppointmentDate.IsZero() {
		input.AppointmentDate = time.Now() // или любое другое значение по умолчанию
	}

	if err := internal.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{Data: input})
}

// GetAppointments godoc
// @Summary Get all appointments
// @Description Get all appointments
// @Tags Appointments
// @Produce  json
// @Success 200 {array} models.Appointment
// @Failure 500 {object} models.ErrorResponse
// @Router /api/admin/appointments [get]
func GetAppointments(c *gin.Context) {
	var appointments []models.Appointment

	// Запрос всех записей из базы данных
	if err := internal.DB.Find(&appointments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{Data: appointments})
}

// ConfirmAppointment godoc
// @Summary Confirm an appointment
// @Description Confirm an appointment by ID
// @Tags Appointments
// @Accept  json
// @Produce  json
// @Param id path uint true "Appointment ID"
// @Success 200 {object} models.Appointment
// @Failure 404 {object} models.ErrorResponse
// @Router /api/admin/appointments/{id}/confirm [put]
func ConfirmAppointment(c *gin.Context) {
	id := c.Param("id")
	var appointment models.Appointment
	if err := internal.DB.First(&appointment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Appointment not found"})
		return
	}

	appointment.Status = "Confirmed"
	internal.DB.Save(&appointment)
	c.JSON(http.StatusOK, models.SuccessResponse{Data: appointment})
}

// DeleteAppointment godoc
// @Summary Delete an appointment
// @Description Delete an appointment by ID
// @Tags Appointments
// @Accept  json
// @Produce  json
// @Param id path uint true "Appointment ID"
// @Success 200 {object} models.SuccessResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /api/admin/appointments/{id} [delete]
func DeleteAppointment(c *gin.Context) {
	id := c.Param("id")
	var appointment models.Appointment
	if err := internal.DB.First(&appointment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Appointment not found"})
		return
	}

	internal.DB.Delete(&appointment)
	c.JSON(http.StatusOK, models.SuccessResponse{Data: true})
}

// CheckAvailability godoc
// @Summary Check appointment availability
// @Description Check if a time slot is available for a specific employee
// @Tags Appointments
// @Accept  json
// @Produce  json
// @Param input body struct{ EmployeeID uint `json:"employee_id" binding:"required"`; DateTime time.Time `json:"date_time" binding:"required"` } true "Availability info"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 409 {object} models.ErrorResponse
// @Router /api/appointments/availability [get]
func CheckAvailability(c *gin.Context) {
	var input struct {
		EmployeeID uint      `json:"employee_id" binding:"required"`
		DateTime   time.Time `json:"date_time" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	// Проверяем наличие записей на указанное время для данного сотрудника
	var count int64
	internal.DB.Model(&models.Appointment{}).Where("employee_id = ? AND appointment_date = ?", input.EmployeeID, input.DateTime).Count(&count)

	if count > 0 {
		c.JSON(http.StatusConflict, models.ErrorResponse{Error: "Time slot not available"})
	} else {
		c.JSON(http.StatusOK, models.SuccessResponse{Data: "Time slot available"})
	}
}
