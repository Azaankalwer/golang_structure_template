package routes

import (
    "github.com/gin-gonic/gin"
    "gin-app/controllers"
)

func UserRoutes(r *gin.Engine) {
    userGroup := r.Group("/users")
    {
        userGroup.GET("/hello", controllers.SayHello)
    }
}
