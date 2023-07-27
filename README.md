# OD-API Documentation

The Buoy API allows you to manage buoys and their associated waves data. Buoys can be created, updated, retrieved, and deleted using this API. Waves data can also be added to specific buoys.

## Endpoints

### Create Buoy

- **URL:** `/buoy`
- **Method:** POST
- **Description:** Create a new buoy.
- **Request Body:**

```json
{
  "buoyname": "Mavericks Buoy",
  "location": "California, USA",
  "payloadType": "waves",
  "batteryVoltage": 4.07,
  "batteryPower": -0.41,
  "solarVoltage": 0.0,
  "humidity": 32.8,
  "waves": [
    {
      "significantWaveHeight": 1.14,
      "peakPeriod": 9.3,
      "meanPeriod": 8.3,
      "peakDirection": 302.3,
      "peakDirectionalSpread": 42.11,
      "meanDirection": 286.2,
      "meanDirectionalSpread": 56.16,
      "timestamp": "2017-11-08T07:06:57.000Z",
      "latitude": 34.30115,
      "longitude": -120.6133
    },
    {
      "significantWaveHeight": 1.14,
      "peakPeriod": 10.24,
      "meanPeriod": 8.44,
      "peakDirection": 312.28,
      "peakDirectionalSpread": 37.07,
      "meanDirection": 284.18,
      "meanDirectionalSpread": 57.4,
      "timestamp": "2017-11-08T07:36:57.000Z",
      "latitude": 34.29883,
      "longitude": -120.61127
    }
  ]
}
```

- **Response:**

```json
{
  "message": "Buoy created successfully",
  "data": "<new_buoy_id>"
}
```

### Get a Buoy

- **URL:** `/buoy/:buoyId`
- **Method:** GET
- **Description:** Retrieve a specific buoy by its ID.
- **Parameters:**
  - `buoyId` (path parameter) - The ID of the buoy to retrieve.
- **Response:**

```json
{
  "data": {
    "id": "<buoy_id>",
    "buoyname": "Mavericks Buoy",
    "location": "California, USA",
    "payloadType": "waves",
    "batteryVoltage": 4.07,
    "batteryPower": -0.41,
    "solarVoltage": 0.0,
    "humidity": 32.8,
    "waves": [
      {
        "significantWaveHeight": 1.14,
        "peakPeriod": 9.3,
        "meanPeriod": 8.3,
        "peakDirection": 302.3,
        "peakDirectionalSpread": 42.11,
        "meanDirection": 286.2,
        "meanDirectionalSpread": 56.16,
        "timestamp": "2017-11-08T07:06:57.000Z",
        "latitude": 34.30115,
        "longitude": -120.6133
      },
      {
        "significantWaveHeight": 1.14,
        "peakPeriod": 10.24,
        "meanPeriod": 8.44,
        "peakDirection": 312.28,
        "peakDirectionalSpread": 37.07,
        "meanDirection": 284.18,
        "meanDirectionalSpread": 57.4,
        "timestamp": "2017-11-08T07:36:57.000Z",
        "latitude": 34.29883,
        "longitude": -120.61127
      }
    ]
  }
}
```

### Update a Buoy

- **URL:** `/buoy/:buoyId`
- **Method:** PUT
- **Description:** Update a specific buoy by its ID.
- **Parameters:**
  - `buoyId` (path parameter) - The ID of the buoy to update.
- **Request Body:**

```json
{
  "buoyname": "New Buoy Name",
  "location": "Updated Location",
  "payloadType": "new_payload",
  "batteryVoltage": 5.0,
  "batteryPower": -0.5,
  "solarVoltage": 1.5,
  "humidity": 40.0
}
```

- **Response:**

```json
{
  "message": "Buoy updated successfully",
  "data": {
    "id": "<buoy_id>",
    "buoyname

": "New Buoy Name",
    "location": "Updated Location",
    "payloadType": "new_payload",
    "batteryVoltage": 5.0,
    "batteryPower": -0.5,
    "solarVoltage": 1.5,
    "humidity": 40.0,
    "waves": [
      {
        "significantWaveHeight": 1.14,
        "peakPeriod": 9.3,
        "meanPeriod": 8.3,
        "peakDirection": 302.3,
        "peakDirectionalSpread": 42.11,
        "meanDirection": 286.2,
        "meanDirectionalSpread": 56.16,
        "timestamp": "2017-11-08T07:06:57.000Z",
        "latitude": 34.30115,
        "longitude": -120.6133
      },
      {
        "significantWaveHeight": 1.14,
        "peakPeriod": 10.24,
        "meanPeriod": 8.44,
        "peakDirection": 312.28,
        "peakDirectionalSpread": 37.07,
        "meanDirection": 284.18,
        "meanDirectionalSpread": 57.4,
        "timestamp": "2017-11-08T07:36:57.000Z",
        "latitude": 34.29883,
        "longitude": -120.61127
      }
    ]
  }
}
```

### Delete a Buoy

- **URL:** `/buoy/:buoyId`
- **Method:** DELETE
- **Description:** Delete a specific buoy by its ID.
- **Parameters:**
  - `buoyId` (path parameter) - The ID of the buoy to delete.
- **Response:**

```json
{
  "message": "Buoy successfully deleted"
}
```

### Get All Buoys

- **URL:** `/buoys`
- **Method:** GET
- **Description:** Retrieve all buoys.
- **Response:**

```json
{
  "data": [
    {
      "id": "<buoy_id_1>",
      "buoyname": "Buoy 1",
      "location": "Location 1",
      "payloadType": "payload_1",
      "batteryVoltage": 4.0,
      "batteryPower": -0.4,
      "solarVoltage": 0.0,
      "humidity": 30.0,
      "waves": [
        {
          "significantWaveHeight": 1.1,
          "peakPeriod": 9.0,
          "meanPeriod": 8.0,
          "peakDirection": 300.0,
          "peakDirectionalSpread": 40.0,
          "meanDirection": 280.0,
          "meanDirectionalSpread": 50.0,
          "timestamp": "2017-11-08T07:00:00.000Z",
          "latitude": 34.0,
          "longitude": -120.0
        }
      ]
    },
    {
      "id": "<buoy_id_2>",
      "buoyname": "Buoy 2",
      "location": "Location 2",
      "payloadType": "payload_2",
      "batteryVoltage": 4.5,
      "batteryPower": -0.5,
      "solarVoltage": 1.5,
      "humidity": 35.0,
      "waves": [
        {
          "significantWaveHeight": 1.2,
          "peakPeriod": 10.0,
          "meanPeriod": 9.0,
          "peakDirection": 310.0,
          "peakDirectionalSpread": 45.0,
          "meanDirection": 290.0,
          "meanDirectionalSpread": 55.0,
          "timestamp": "2017-11-08T07:15:00.000Z",
          "latitude": 35.0,
          "longitude": -121.0
        }
      ]
    }
  ]
}
```

### Add Waves Data to a Buoy

- **URL:** `/buoy/:buoyId/waves`
- **Method:** POST
- **Description:** Add waves data to a specific buoy by its ID.
- **Parameters:**
  - `buoyId` (path parameter) - The ID of the buoy to add waves data to.
- **Request Body:**

```json
{
  "significantWaveHeight": 1.3,
  "peakPeriod": 10.5,
  "meanPeriod": 9.5,
  "peakDirection": 315.0,
  "peakDirectionalSpread": 50.0,
  "meanDirection": 295.0,
  "meanDirectionalSpread": 60.0,
  "timestamp": "2017-11-08T07:30:00.000Z",
  "latitude": 36.0,
  "longitude": -122.0
}
```

- **Response:**

```json
{
  "message": "Waves data added to buoy successfully"
}
```

## Error Responses

In case of errors, the API will respond with appropriate error messages and status codes. Here are some possible error responses:

### Example: Bad Request

```json
{
  "error": "Invalid request"
}
```

### Example: Not Found

```json
{
  "error": "Buoy not found"
}
```

### Example: Internal Server Error

```json
{
  "error": "Failed to create buoy"
}
```

## Conclusion

This Buoy API documentation provides information on how to interact with the API to manage buoys and their waves data. It covers all the available endpoints, request and response formats, and possible error responses. Feel free to use this API to manage your buoy data effectively!