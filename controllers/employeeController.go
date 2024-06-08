package controllers

import (
	"autoservice-backend/internal"
	"autoservice-backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetEmployees godoc
// @Summary Get all employees
// @Description Get all employees
// @Tags Employees
// @Produce json
// @Param role query string false "Role"
// @Success 200 {array} models.Employee
// @Failure 500 {object} models.ErrorResponse
// @Router /api/home [get]
func GetEmployees(c *gin.Context) {
	var employees []models.Employee
	role := c.Query("role")
	query := internal.DB
	if role != "" {
		query = query.Where("role = ?", role)
	}
	query.Find(&employees)
	c.JSON(http.StatusOK, models.SuccessResponse{Data: employees})
}

// CreateEmployee godoc
// @Summary Create a new employee
// @Description Create a new employee
// @Tags Employees
// @Accept  json
// @Produce  json
// @Param input body models.Employee true "Employee info"
// @Success 200 {object} models.Employee
// @Failure 400 {object} models.ErrorResponse
// @Router /api/admin/employees [post]
func CreateEmployee(c *gin.Context) {
	var input models.Employee
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}
	internal.DB.Create(&input)
	c.JSON(http.StatusOK, models.SuccessResponse{Data: input})
}

// UpdateEmployee godoc
// @Summary Update an employee
// @Description Update an employee by ID
// @Tags Employees
// @Accept  json
// @Produce  json
// @Param id path uint true "Employee ID"
// @Param input body models.Employee true "Employee info"
// @Success 200 {object} models.Employee
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /api/admin/employees/{id} [put]
func UpdateEmployee(c *gin.Context) {
	id := c.Param("id")
	var employee models.Employee
	if err := internal.DB.First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Employee not found"})
		return
	}

	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	internal.DB.Save(&employee)
	c.JSON(http.StatusOK, models.SuccessResponse{Data: employee})
}

// DeleteEmployee godoc
// @Summary Delete an employee
// @Description Delete an employee by ID
// @Tags Employees
// @Accept  json
// @Produce  json
// @Param id path uint true "Employee ID"
// @Success 200 {object} models.SuccessResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /api/admin/employees/{id} [delete]
func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")
	var employee models.Employee
	if err := internal.DB.First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Employee not found"})
		return
	}

	internal.DB.Delete(&employee)
	c.JSON(http.StatusOK, models.SuccessResponse{Data: true})
}
