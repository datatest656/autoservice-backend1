package routes

import (
	"autoservice-backend/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Добавьте CORS Middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Разрешите доступ с вашего фронтенда
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Маршрут для документации Swagger
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Маршруты для публичной части
	router.GET("/api/home", controllers.GetEmployees)
	router.GET("/api/services", controllers.GetServices)
	router.GET("/api/services/:id", controllers.GetServiceByID)
	router.POST("/api/appointments", controllers.CreateAppointment)
	router.GET("/api/appointments/availability", controllers.CheckAvailability)
	router.GET("/api/news", controllers.GetNews)
	router.GET("/api/news/:id", controllers.GetNewsByID)
	router.GET("/api/contact", controllers.GetContactInfo)
	router.POST("/api/feedback", controllers.CreateFeedback)

	// Маршруты для административной панели
	adminRoutes := router.Group("/api/admin")

	// Управление сотрудниками
	adminRoutes.POST("/employees", controllers.CreateEmployee)
	adminRoutes.PUT("/employees/:id", controllers.UpdateEmployee)
	adminRoutes.DELETE("/employees/:id", controllers.DeleteEmployee)

	// Управление услугами
	adminRoutes.POST("/services", controllers.CreateService)
	adminRoutes.PUT("/services/:id", controllers.UpdateService)
	adminRoutes.DELETE("/services/:id", controllers.DeleteService)

	// Управление записями
	adminRoutes.GET("/appointments", controllers.GetAppointments)
	adminRoutes.PUT("/appointments/:id/confirm", controllers.ConfirmAppointment)
	adminRoutes.DELETE("/appointments/:id", controllers.DeleteAppointment)

	// Управление новостями и акциями
	adminRoutes.POST("/news", controllers.CreateNews)
	adminRoutes.PUT("/news/:id", controllers.UpdateNews)
	adminRoutes.DELETE("/news/:id", controllers.DeleteNews)

	// Просмотр статистики
	adminRoutes.GET("/statistics", controllers.GetStatistics)
	router.POST("/api/clients", controllers.CreateClient)
	router.GET("/api/clients", controllers.GetClients)

	return router
}
