package routers

import (
    "github.com/gin-gonic/gin"
    "gnz_psi/go_web_service/controllers"
)

// RegisterUserRoutes defines user-related routes
func PSIDataRoutes(router *gin.Engine) {

    userGroup := router.Group("/psi")
    {
        userGroup.GET("/", controllers.GetUserHandler)
        userGroup.POST("/", controllers.CreateUserHandler)
    }

}
