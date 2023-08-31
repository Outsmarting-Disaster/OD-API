package controllers

import (
	"context"
	"net/http"
	"time"
	"math/rand"
	// "fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"od-api/configs"
	"od-api/models"
	"od-api/responses"
)

var buoyCollection *mongo.Collection = configs.GetCollection(configs.DB, "buoys")

func CreateBuoy() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var buoy models.Buoy
		defer cancel()

		// Validate the request body
		if err := c.BindJSON(&buoy); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		// Insert the buoy into the database using the provided MongoDB collection
		result, err := buoyCollection.InsertOne(ctx, buoy)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create buoy"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Buoy created successfully", "data": result.InsertedID})

		// Start a Goroutine to periodically post waves data for this buoy
		// go postWavesDataPeriodically(ctx, buoy)
	}
}

func GetABuoy() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		buoyID := c.Param("buoyId")
		defer cancel()

		objID, err := primitive.ObjectIDFromHex(buoyID)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.BuoyResponse{
				Status:  http.StatusBadRequest,
				Message: "Invalid buoy ID",
				Data:    nil,
			})
			return
		}

		var buoy models.Buoy
		err = buoyCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&buoy)
		if err != nil {
			c.JSON(http.StatusNotFound, responses.BuoyResponse{
				Status:  http.StatusNotFound,
				Message: "Buoy not found",
				Data:    nil,
			})
			return
		}

		c.JSON(http.StatusOK, responses.BuoyResponse{
			Status:  http.StatusOK,
			Message: "Buoy found",
			Data:    map[string]interface{}{"buoy": buoy},
		})
	}
}

func EditBuoy() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		buoyID := c.Param("buoyId")
		var buoy models.Buoy
		defer cancel()

		objID, err := primitive.ObjectIDFromHex(buoyID)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.BuoyResponse{
				Status:  http.StatusBadRequest,
				Message: "Invalid buoy ID",
				Data:    nil,
			})
			return
		}

		// Validate the request body
		if err := c.BindJSON(&buoy); err != nil {
			c.JSON(http.StatusBadRequest, responses.BuoyResponse{
				Status:  http.StatusBadRequest,
				Message: "Invalid request",
				Data:    nil,
			})
			return
		}

		// Update the buoy in the database using the provided MongoDB collection
		update := bson.M{
			"buoyname":       buoy.BuoyName,
			"location":       buoy.Location,
			"payloadType":    buoy.PayloadType,
			"batteryVoltage": buoy.BatteryVoltage,
			"batteryPower":   buoy.BatteryPower,
			"solarVoltage":   buoy.SolarVoltage,
			"humidity":       buoy.Humidity,
			"waves":          buoy.Waves,
		}

		result, err := buoyCollection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BuoyResponse{
				Status:  http.StatusInternalServerError,
				Message: "Failed to update buoy",
				Data:    nil,
			})
			return
		}

		// Get updated buoy details
		var updatedBuoy models.Buoy
		if result.MatchedCount == 1 {
			err := buoyCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&updatedBuoy)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.BuoyResponse{
					Status:  http.StatusInternalServerError,
					Message: "Failed to get updated buoy details",
					Data:    nil,
				})
				return
			}
		}

		c.JSON(http.StatusOK, responses.BuoyResponse{
			Status:  http.StatusOK,
			Message: "Buoy updated successfully",
			Data:    map[string]interface{}{"buoy": updatedBuoy},
		})
	}
}

func DeleteBuoy() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		buoyID := c.Param("buoyId")
		defer cancel()

		objID, err := primitive.ObjectIDFromHex(buoyID)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.BuoyResponse{
				Status:  http.StatusBadRequest,
				Message: "Invalid buoy ID",
				Data:    nil,
			})
			return
		}

		result, err := buoyCollection.DeleteOne(ctx, bson.M{"_id": objID})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BuoyResponse{
				Status:  http.StatusInternalServerError,
				Message: "Failed to delete buoy",
				Data:    nil,
			})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound, responses.BuoyResponse{
				Status:  http.StatusNotFound,
				Message: "Buoy with specified ID not found!",
				Data:    nil,
			})
			return
		}

		c.JSON(http.StatusOK, responses.BuoyResponse{
			Status:  http.StatusOK,
			Message: "Buoy successfully deleted",
			Data:    nil,
		})
	}
}

func GetAllBuoys() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var buoys []models.Buoy
		defer cancel()

		results, err := buoyCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BuoyResponse{
				Status:  http.StatusInternalServerError,
				Message: "Failed to get all buoys",
				Data:    nil,
			})
			return
		}

		// Read from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleBuoy models.Buoy
			if err = results.Decode(&singleBuoy); err != nil {
				c.JSON(http.StatusInternalServerError, responses.BuoyResponse{
					Status:  http.StatusInternalServerError,
					Message: "Failed to decode buoy data",
					Data:    nil,
				})
				return
			}

			buoys = append(buoys, singleBuoy)
		}

		c.JSON(http.StatusOK, responses.BuoyResponse{
			Status:  http.StatusOK,
			Message: "Buoys found",
			Data:    map[string]interface{}{"buoys": buoys},
		})
	}
}
func AddWavesDataToBuoy() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		buoyID := c.Param("buoyId")
		var wavesData models.WavesData
		defer cancel()

		objID, err := primitive.ObjectIDFromHex(buoyID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid buoy ID"})
			return
		}

		// Validate the request body
		if err := c.BindJSON(&wavesData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		// Find the buoy with the specified ID
		var buoy models.Buoy
		err = buoyCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&buoy)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Buoy not found"})
			return
		}

		// Append the new waves data to the buoy's Waves slice
		buoy.Waves = append(buoy.Waves, wavesData)

		// Update the buoy in the database with the new waves data
		update := bson.M{"$set": bson.M{"waves": buoy.Waves}}
		_, err = buoyCollection.UpdateOne(ctx, bson.M{"_id": objID}, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add waves data to buoy"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Waves data added to buoy successfully"})
	}
}



const (
	MinLatitude  = -90.0
	MaxLatitude  = 90.0
	MinLongitude = -180.0
	MaxLongitude = 180.0
)
// Generate realistic latitude and longitude for the buoy's location
func generateRandomLocation(initialLatitude, initialLongitude float64) (float64, float64) {
	// Calculate random deviation from the initial latitude and longitude
	latitudeDeviation := rand.Float64() * 0.01
	longitudeDeviation := rand.Float64() * 0.01

	// Calculate the new latitude and longitude by adding the deviation to the initial coordinates
	latitude := initialLatitude + latitudeDeviation
	longitude := initialLongitude + longitudeDeviation

	// Clamp latitude and longitude to the valid range
	if latitude < MinLatitude {
		latitude = MinLatitude
	} else if latitude > MaxLatitude {
		latitude = MaxLatitude
	}
	if longitude < MinLongitude {
		longitude = MinLongitude
	} else if longitude > MaxLongitude {
		longitude = MaxLongitude
	}

	return latitude, longitude
}

// Generate realistic wave data
func generateRandomWavesData(initialLatitude, initialLongitude float64) models.WavesData {
	// Generate random wave height (between 0.5 and 5 meters)
	significantWaveHeight := rand.Float64()*4.5 + 0.5

	// Generate random wave period (between 4 and 15 seconds)
	peakPeriod := rand.Float64()*11.0 + 4.0

	// Generate random wave direction (between 0 and 360 degrees)
	peakDirection := rand.Float64() * 360.0

	// Generate random timestamp within the last 24 hours
	timestamp := time.Now().Add(-time.Duration(rand.Intn(24))*time.Hour).UTC().Format(time.RFC3339)

	// Generate random latitude and longitude for the buoy's location
	latitude, longitude := generateRandomLocation(initialLatitude, initialLongitude)

	// Calculate random direction for the waves (deviation from peakDirection)
	peakDirectionalSpread := rand.Float64()*30.0 + 5.0

	// Calculate mean direction (opposite to peakDirection)
	meanDirection := peakDirection + 180.0
	if meanDirection > 360.0 {
		meanDirection -= 360.0
	}

	// Calculate random spread for the mean direction
	meanDirectionalSpread := rand.Float64()*60.0 + 15.0

	return models.WavesData{
		SignificantWaveHeight:   significantWaveHeight,
		PeakPeriod:              peakPeriod,
		MeanPeriod:              peakPeriod * 0.9,
		PeakDirection:           peakDirection,
		PeakDirectionalSpread:   peakDirectionalSpread,
		MeanDirection:           meanDirection,
		MeanDirectionalSpread:   meanDirectionalSpread,
		Timestamp:               timestamp,
		Latitude:                latitude,
		Longitude:               longitude,
	}
}

// Insert wave data into the database for a specific buoy ID
func InsertWaveDataForBuoy(buoyID string, waveData models.WavesData) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(buoyID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$push": bson.M{"waves": waveData}}

	_, err = buoyCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func CreateWaveDataForBuoy() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data models.WavesData

		// Validate the request body
		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "Invalid request",
				Data:    map[string]interface{}{},
			})
			return
		}

		// You can retrieve the buoyID from the request, whether it's from a URL parameter or JSON data
		buoyID := "64c1de1bccc77c103ab51ed1" // Replace this with the actual buoy ID from the request

		// Insert the wave data for the specified buoy ID
		if err := InsertWaveDataForBuoy(buoyID, data); err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "Failed to insert wave data",
				Data:    map[string]interface{}{},
			})
			return
		}

		c.JSON(http.StatusCreated, responses.UserResponse{
			Status:  http.StatusCreated,
			Message: "Wave data created successfully",
			Data:    map[string]interface{}{},
		})
	}
}