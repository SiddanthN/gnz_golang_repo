package routers

import (
    "github.com/gin-gonic/gin"
    "gnz_psi/go_web_service/controllers"
)

// PSIDataRouter defines PSI Data related routes
func PSIDataRouter(router *gin.Engine) {

    PSIRouteGroup := router.Group("/gnz/api")
    {
        apiVersion := PSIRouteGroup.Group("/v1")
        {
            action := apiVersion.Group("/psi")
            {
                action.GET("", controllers.GetPSIData)       // GET /gnz/api/v1/psi
                action.POST("", controllers.CreateUserHandler) // POST /gnz/api/v1/psi
            }

        }
    }

}
