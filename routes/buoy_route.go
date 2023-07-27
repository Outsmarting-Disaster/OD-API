package routes

import (
	"od-api/controllers"
    "github.com/gin-gonic/gin"
)

func BuoyRoute(router *gin.Engine) {
	router.POST("/buoy", controllers.CreateBuoy())
	router.GET("/buoy/:buoyId", controllers.GetABuoy())
	router.PUT("/buoy/:buoyId", controllers.EditBuoy())
	router.DELETE("/buoy/:buoyId", controllers.DeleteBuoy())
	router.GET("/buoys", controllers.GetAllBuoys())
	router.POST("/buoy/:buoyId/waves", controllers.AddWavesDataToBuoy()) // New endpoint to add waves data
}