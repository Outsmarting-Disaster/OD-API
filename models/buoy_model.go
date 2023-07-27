package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type WavesData struct {
	SignificantWaveHeight   float64 `json:"significantWaveHeight"`
	PeakPeriod              float64 `json:"peakPeriod"`
	MeanPeriod              float64 `json:"meanPeriod"`
	PeakDirection           float64 `json:"peakDirection"`
	PeakDirectionalSpread   float64 `json:"peakDirectionalSpread"`
	MeanDirection           float64 `json:"meanDirection"`
	MeanDirectionalSpread   float64 `json:"meanDirectionalSpread"`
	Timestamp               string  `json:"timestamp"`
	Latitude                float64 `json:"latitude"`
	Longitude               float64 `json:"longitude"`
}

type Buoy struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	BuoyName       string             `json:"buoyname,omitempty" validate:"required"`
	Location       string             `json:"location,omitempty" validate:"required"`
	PayloadType    string             `json:"payloadType,omitempty" validate:"required"`
	BatteryVoltage float64            `json:"batteryVoltage,omitempty"`
	BatteryPower   float64            `json:"batteryPower,omitempty"`
	SolarVoltage   float64            `json:"solarVoltage,omitempty"`
	Humidity       float64            `json:"humidity,omitempty"`
	Waves          []WavesData        `json:"waves,omitempty"`
}
