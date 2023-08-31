package main

import (
        "time"
	"fmt"

	"od-api/configs"
	"od-api/routes" //add this
        "od-api/controllers"
	"github.com/gin-gonic/gin"
)

var buoyInitialCoordinates = map[string]struct {
	InitialLatitude  float64
	InitialLongitude float64
}{
	"64c1de1bccc77c103ab51ed1": {InitialLatitude: 34.30115, InitialLongitude: -120.6133},
	"6a3dcff47db9bc1c9e8d5c84": {InitialLatitude: 34.29883, InitialLongitude: -120.61127},
	"4a90e12bd287d0914d29b197": {InitialLatitude: 34.30212, InitialLongitude: -120.61201},
	"1b2fd03771ad3a8abf25f82e": {InitialLatitude: 34.30027, InitialLongitude: -120.60915},
}

func generateAndInsertData(buoyID string, initialLatitude, initialLongitude float64) {
	for {
		// Generate realistic wave data for the specific buoy
		waveData := controllers.GenerateRandomWavesData(initialLatitude, initialLongitude)

		// Insert wave data into the database for the specific buoy ID
		err := controllers.InsertWaveDataForBuoy(buoyID, waveData)
		if err != nil {
			fmt.Println("Failed to insert wave data for buoy", buoyID, ":", err)
		}

		// Wait for 1 minute before generating and inserting the next data for this buoy
		time.Sleep(1 * time.Minute)
	}
}

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
         go generateAndInsertData("64c1de1bccc77c103ab51ed1", 34.30115, -120.6133) // Replace with the actual buoy ID and initial latitude/longitude
	// go generateAndInsertData("your-buoy-id-2", your-lat-2, your-long-2) // Repeat this line for other buoys
        router.Run("localhost:6000") 
}