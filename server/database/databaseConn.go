package database

import (
	"log"
	"os"

	"github.com/shiibs/nosh_dashboard/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DBConn is a global variable to hold the database connection.
var DBConn *gorm.DB

// ConnectDB initializes the database connection using the DATABASE_URL environment variable.
func ConnectDB() {
	// Get the database URL from the environment variables.
	dsn := os.Getenv("DATABASE_URL")

	// Open a connection to the PostgreSQL database using GORM.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Set the logger to log errors only.
		Logger: logger.Default.LogMode(logger.Error),
	})

	// If there is an error opening the connection, panic and stop execution.
	if err != nil {
		panic("Database connection failed")
	}

	log.Println("DB connected")

	// Automatically migrate the schema for the Dish model.
	if err = db.AutoMigrate(new(models.Dish)); err != nil {
		log.Println(err)
	}

	// Assign the database connection to the global variable.
	DBConn = db
}
