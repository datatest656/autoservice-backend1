package controllers

import (
	"autoservice-backend/internal"
	"autoservice-backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetNews godoc
// @Summary Get all news
// @Description Get all news
// @Tags News
// @Produce json
// @Success 200 {array} models.News
// @Failure 500 {object} models.ErrorResponse
// @Router /api/news [get]
func GetNews(c *gin.Context) {
	var news []models.News
	internal.DB.Find(&news)
	c.JSON(http.StatusOK, models.SuccessResponse{Data: news})
}

// GetNewsByID godoc
// @Summary Get news by ID
// @Description Get news by ID
// @Tags News
// @Produce json
// @Param id path uint true "News ID"
// @Success 200 {object} models.News
// @Failure 404 {object} models.ErrorResponse
// @Router /api/news/{id} [get]
func GetNewsByID(c *gin.Context) {
	id := c.Param("id")
	var news models.News
	if err := internal.DB.First(&news, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "News not found"})
		return
	}
	c.JSON(http.StatusOK, models.SuccessResponse{Data: news})
}

// CreateNews godoc
// @Summary Create news
// @Description Create a new news
// @Tags News
// @Accept  json
// @Produce  json
// @Param input body models.News true "News info"
// @Success 200 {object} models.News
// @Failure 400 {object} models.ErrorResponse
// @Router /api/admin/news [post]
func CreateNews(c *gin.Context) {
	var input models.News
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}
	internal.DB.Create(&input)
	c.JSON(http.StatusOK, models.SuccessResponse{Data: input})
}

// UpdateNews godoc
// @Summary Update news
// @Description Update news by ID
// @Tags News
// @Accept  json
// @Produce  json
// @Param id path uint true "News ID"
// @Param input body models.News true "News info"
// @Success 200 {object} models.News
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /api/admin/news/{id} [put]
func UpdateNews(c *gin.Context) {
	id := c.Param("id")
	var news models.News
	if err := internal.DB.First(&news, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "News not found"})
		return
	}

	if err := c.ShouldBindJSON(&news); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	internal.DB.Save(&news)
	c.JSON(http.StatusOK, models.SuccessResponse{Data: news})
}

// DeleteNews godoc
// @Summary Delete news
// @Description Delete news by ID
// @Tags News
// @Accept  json
// @Produce  json
// @Param id path uint true "News ID"
// @Success 200 {object} models.SuccessResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /api/admin/news/{id} [delete]
func DeleteNews(c *gin.Context) {
	id := c.Param("id")
	var news models.News
	if err := internal.DB.First(&news, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "News not found"})
		return
	}

	internal.DB.Delete(&news)
	c.JSON(http.StatusOK, models.SuccessResponse{Data: true})
}
