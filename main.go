package main

import (
	"fmt"
	"trainingbackenddot/config"
	"trainingbackenddot/infrastructure/db"
	"trainingbackenddot/interface/http"
	"trainingbackenddot/router"
	"trainingbackenddot/usecase"
)

func main() {
	// Connection to Database
	config.ConnectDatabase()
	db.MigrateDatabase()

	// Repository Initialization
	adminRepo := db.NewAdminRepository(config.DB)
	userRepo := db.NewUserRepository(config.DB)
	studioRepo := db.NewStudioRepository(config.DB)
	filmRepo := db.NewFilmRepository(config.DB)
	scheduleRepo := db.NewScheduleRepository(config.DB)
	ticketRepo := db.NewTicketRepository(config.DB)
	transactionRepo := db.NewTransactionRepository(config.DB)
	notificationRepo := db.NewNotificationRepository(config.DB)

	// UseCase Initialization
	adminUC := usecase.NewAdminUseCase(adminRepo)
	userUC := usecase.NewUserUseCase(userRepo)
	studioUC := usecase.NewStudioUsecase(studioRepo)
	filmUC := usecase.NewFilmUsecase(filmRepo)
	scheduleUC := usecase.NewScheduleUsecase(scheduleRepo)
	ticketUC := usecase.NewTicketUsecase(ticketRepo, scheduleRepo)
	transactionUC := usecase.NewTransactionUsecase(transactionRepo, ticketRepo, scheduleRepo, studioRepo)
	notificationUC := usecase.NewNotificationUsecase(notificationRepo)

	// Handler Initialization
	adminHandler := http.NewAdminHandler(adminUC)
	userHandler := http.NewUserHandler(userUC)
	studioHandler := http.NewStudioHandler(studioUC)
	filmHandler := http.NewFilmHandler(filmUC)
	scheduleHandler := http.NewScheduleHandler(scheduleUC, notificationUC)
	ticketHandler := http.NewTicketHandler(ticketUC)
	transactionHandler := http.NewTransactionHandler(transactionUC)
	notificationHandler := http.NewNotificationHandler(notificationUC)

	// Router Setup
	r := router.SetupRouter(adminHandler, userHandler, studioHandler, filmHandler, scheduleHandler, ticketHandler, transactionHandler, notificationHandler)

	// Running the Server
	fmt.Println("The server is running on http://localhost:3000")
	r.Run(":3000")
}
