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

	// UseCase Initialization
	adminUC := usecase.NewAdminUseCase(adminRepo)
	userUC := usecase.NewUserUseCase(userRepo)
	studioUC := usecase.NewStudioUsecase(studioRepo)
	filmUC := usecase.NewFilmUsecase(filmRepo)

	// Handler Initialization
	adminHandler := http.NewAdminHandler(adminUC)
	userHandler := http.NewUserHandler(userUC)
	studioHandler := http.NewStudioHandler(studioUC)
	filmHandler := http.NewFilmHandler(filmUC)

	// Router Setup
	r := router.SetupRouter(adminHandler, userHandler, studioHandler, filmHandler)

	// Running the Server
	fmt.Println("The server is running on http://localhost:3000")
	r.Run(":3000")
}
