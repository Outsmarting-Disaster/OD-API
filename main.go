package main

import (
	"od-api/configs"
	"od-api/routes" //add this
	"github.com/gin-gonic/gin"
)

func main() {
        router := gin.Default()

        router.GET("/", func(c *gin.Context) {
                c.JSON(200, gin.H{
                        "data": "Hello from Gin-gonic & mongoDB",
                })
        })

		// run database
		configs.ConnectDB()

		routes.UserRoute(router) //add this
                routes.BuoyRoute(router)

        router.Run("localhost:6000") 
}