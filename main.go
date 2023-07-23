package main

import (
	"outsmarting-disaster-api/configs"
	"outsmarting-disaster-api/routes" //add this
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

        router.Run("localhost:6000") 
}