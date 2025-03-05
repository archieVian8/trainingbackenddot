package router

import (
	"trainingbackenddot/interface/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	adminHandler *http.AdminHandler,
	userHandler *http.UserHandler,
	studioHandler *http.StudioHandler,
	filmHandler *http.FilmHandler,
	scheduleHandler *http.ScheduleHandler,
	ticketHandler *http.TicketHandler,
	transactionHandler *http.TransactionHandler) *gin.Engine {
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

			// Film Management (Admin Only)
			film := admin.Group("/films")
			{
				film.POST("", filmHandler.AddFilm)
				film.PUT("/:id", filmHandler.UpdateFilm)
				film.DELETE("/:id", filmHandler.DeleteFilm)
				film.GET("/viewall", filmHandler.GetAllFilms)
			}

			// Movie Schedule Management (Admin Only)
			schedule := admin.Group("/schedules")
			{
				schedule.POST("/", scheduleHandler.CreateSchedule)
				schedule.PUT("/:id", scheduleHandler.UpdateSchedule)
				schedule.DELETE("/:id", scheduleHandler.DeleteSchedule)
				schedule.GET("/viewall", scheduleHandler.ViewAllSchedules)
				schedule.POST("promo/:id", scheduleHandler.ApplyPromo)
			}

			// View All Transactions (Admin Only)
			transactions := admin.Group("/transactions")
			{
				transactions.GET("/viewall", transactionHandler.ViewAllTransactions)
				transactions.GET("/viewfilm/daily", transactionHandler.ViewDailySalesByFilm)
				transactions.GET("/viewfilm/monthly", transactionHandler.ViewMonthlySalesByFilm)
				transactions.GET("/viewstudio/daily", transactionHandler.ViewDailySalesByStudio)
				transactions.GET("/viewstudio/monthly", transactionHandler.ViewMonthlySalesByStudio)

			}

		}

		user := api.Group("/user")
		{
			// User Auth
			user.POST("/signup", userHandler.SignupUser)
			user.POST("/signin", userHandler.SigninUser)
			user.GET("/viewall", userHandler.ViewAllUsers)

			// User view Schedule
			schedule := user.Group("/schedules")
			{
				schedule.GET("/viewall", scheduleHandler.ViewAllSchedules)
			}

			// Ticketing
			ticket := user.Group("/tickets")
			{
				ticket.POST("/book", ticketHandler.BookTicket)
				ticket.GET("/:id", ticketHandler.GetTicket)
			}

			// Transaction
			transaction := user.Group("/transactions")
			{
				transaction.POST("/pay/:id", transactionHandler.PayTicket)
			}
		}
	}

	return router
}
