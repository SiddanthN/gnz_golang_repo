package main

import (
    "log"
    "gnz_psi/go_web_service/routers"
)

func main() {
    router := routers.SetupRouter()

    log.Println("Server started at :8080")
    router.Run(":8080") // Gin's built-in HTTP server
}
