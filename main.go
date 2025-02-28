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

	// UseCase Initialization
	adminUC := usecase.NewAdminUseCase(adminRepo)
	userUC := usecase.NewUserUseCase(userRepo)

	// Handler Initialization
	adminHandler := http.NewAdminHandler(adminUC)
	userHandler := http.NewUserHandler(userUC)

	// Router Setup
	r := router.SetupRouter(adminHandler, userHandler)

	// Running the Server
	fmt.Println("The server is running on http://localhost:3000")
	r.Run(":3000")
}
