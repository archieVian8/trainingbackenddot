package router

import (
	"trainingbackenddot/interface/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	adminHandler *http.AdminHandler,
	userHandler *http.UserHandler,
	studioHandler *http.StudioHandler) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		admin := api.Group("/admin")
		{
			// Admin Auth
			admin.POST("/signup", adminHandler.SignupAdmin)
			admin.POST("/signin", adminHandler.SigninAdmin)
			admin.GET("/viewall", adminHandler.ViewAllAdmins)

			// Studio Management (Admin Only)
			studios := admin.Group("/studios")
			{
				studios.POST("", studioHandler.CreateStudio)
				studios.PUT("/:id", studioHandler.UpdateStudio)
				studios.DELETE("/:id", studioHandler.DeleteStudio)
				studios.GET("/viewall", studioHandler.GetAllStudios)
			}
		}

		user := api.Group("/user")
		{
			// Admin Auth
			user.POST("/signup", userHandler.SignupUser)
			user.POST("/signin", userHandler.SigninUser)
			user.GET("/viewall", userHandler.ViewAllUsers)
		}
	}

	return router
}
