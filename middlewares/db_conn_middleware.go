package middlewares

import (

	"os"
	"log"
	"fmt"
	"errors"
	"syscall"
	"os/signal"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/driver/sqlite"

	"gnz_psi/go_web_service/configs"

)

// Global DB instance
var DatabaseInterface *gorm.DB

// Initialize database connection
func InitDB() {

	fmt.Println("[DEBUG] Database file path: ", configs.SqliteDBFilename)

	var databaseConnErr error
	DatabaseInterface, databaseConnErr = gorm.Open(sqlite.Open(configs.SqliteDBFilename), &gorm.Config{

			NamingStrategy: schema.NamingStrategy{

				SingularTable: true,

			},
	
		},

	)

	if databaseConnErr != nil {

		log.Fatal("Failed to connect to database:", databaseConnErr)

	}

	log.Println("Database connection initialized")

}

// GetDB returns the global DB instance
func GetDBInterface() (*gorm.DB, error) {

    if DatabaseInterface == nil {

        return nil, errors.New("Error: Database connection could not be initialized")

    }

	return DatabaseInterface, nil

}

// CloseDB closes the database connection
func CloseDB() {

	if DatabaseInterface != nil {

		DatabaseConn, err := DatabaseInterface.DB()

		if err != nil {

			log.Println("Failed to get SQLite DB instance:", err)
			return

		}

		DatabaseConn.Close()

		log.Println("Database connection closed")

	}

}

// GracefulShutdown listens for OS signals and cleans up resources
func GracefulShutdown() {

	quit := make(chan os.Signal, 1)
	
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit // Capture SIGINT (Ctrl+C) & SIGTERM

	log.Println("Shutting database connection down gracefully...")

	CloseDB()

	os.Exit(0)

}
