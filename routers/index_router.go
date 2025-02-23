package routers

import (
	"gnz_psi/go_web_service/middlewares"
    "github.com/gin-gonic/gin"
)

// Setting up the main router for the service
func SetupRouter() *gin.Engine {

    router := gin.Default()

	// Applying middleware to all the routers here
	router.Use(middlewares.LoggingMiddleware())

    // Lisitng down all the group routers
    PSIDataRouter(router)
	
    return router

}
