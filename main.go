
package main

import (
    "os"
    "fmt"
	"log"
    "github.com/joho/godotenv"
	"gnz_psi/go_web_service/routers"
	"gnz_psi/go_web_service/middlewares"
)

func loadEnv() {

	if err := godotenv.Load(".env"); err != nil {

		log.Println("Warning: Could not load .env file. Terminating server boot up.")
		os.Exit(1)

	}

	log.Println("Environment file loaded successfully.")

}

func main() {

    // Loading the environment variables
    loadEnv();

	// Initialize the database
	middlewares.InitDB()

	// Handle graceful shutdown in a separate goroutine
	go middlewares.GracefulShutdown()

	// Start the router
	router := routers.SetupRouter()

    dotenv_load_err := godotenv.Load(".env")

    if dotenv_load_err != nil{

        log.Fatalf("Error loading .env file: %s", dotenv_load_err)

    } else {

        server_port:= os.Getenv("SERVER_PORT")
        

        log.Println(fmt.Sprintf("Server up and running at: %s on port: %s", os.Getenv("SERVER_HOST"), server_port))
        router.Run(fmt.Sprintf(":%s", server_port))

    }

}
