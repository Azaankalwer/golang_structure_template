package main

import (
    "gin-app/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
	
    routes.UserRoutes(r)

    r.Run(":8080")
}
