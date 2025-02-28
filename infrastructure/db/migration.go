package db

import (
	"fmt"
	"log"
	"trainingbackenddot/config"
	"trainingbackenddot/domain"
)

func MigrateDatabase() {
	if config.DB == nil {
		log.Fatal("The database is not connected! Run ConnectDatabase() first.")
	}

	err := config.DB.AutoMigrate(
		&domain.Admin{},
		&domain.User{},
		&domain.Studio{},
		&domain.Film{},
		&domain.Schedule{},
		&domain.Ticket{},
		&domain.Transaction{},
	)
	if err != nil {
		log.Fatal("Failed to migrate:", err)
	}

	fmt.Println("Database migration successful")
}
