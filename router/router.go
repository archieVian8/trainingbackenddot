package router

import (
	"trainingbackenddot/interface/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(adminHandler *http.AdminHandler, userHandler *http.UserHandler) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		admin := api.Group("/admin")
		{
			admin.POST("/signup", adminHandler.SignupAdmin)
			admin.POST("/signin", adminHandler.SigninAdmin)
			admin.GET("/viewall", adminHandler.ViewAllAdmins)
		}

		user := api.Group("/user")
		{
			user.POST("/signup", userHandler.SignupUser)
			user.POST("/signin", userHandler.SigninUser)
			user.GET("/viewall", userHandler.ViewAllUsers)
		}
	}

	return router
}
