package controllers

import (
	"context"
	"net/http"
	"time"

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
			c.JSON(http.StatusBadRequest, responses.BuoyResponse{
				Status:  http.StatusBadRequest,
				Message: "Invalid request",
				Data:    nil,
			})
			return
		}

		// Insert the buoy into the database using the provided MongoDB collection
		result, err := buoyCollection.InsertOne(ctx, buoy)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BuoyResponse{
				Status:  http.StatusInternalServerError,
				Message: "Failed to create buoy",
				Data:    nil,
			})
			return
		}

		c.JSON(http.StatusCreated, responses.BuoyResponse{
			Status:  http.StatusCreated,
			Message: "Buoy created successfully",
			Data:    map[string]interface{}{"id": result.InsertedID},
		})
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
