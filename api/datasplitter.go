package api

import "wim-api/domain"

func bulkToSimple(data domain.RawInputData) (domain.VehicleData, domain.Coordinate) {
	var vehicleData domain.VehicleData
	var coordinate domain.Coordinate

	vehicleData.ID = data.ID
	vehicleData.VehicleSpeed = data.VehicleSpeed
	vehicleData.Acceleration = data.Acceleration
	vehicleData.CoolentTemp = data.CoolentTemp
	vehicleData.ElevationAngle = data.ElevationAngle
	vehicleData.EngineLoad = data.EngineLoad
	vehicleData.FuelFlow = data.FuelFlow
	vehicleData.HeadingDirection = data.HeadingDirection
	vehicleData.IntakeAirTemp = data.IntakeAirTemp
	vehicleData.O2 = data.O2
	vehicleData.Rpm = data.Rpm
	vehicleData.Weight = data.Weight
	vehicleData.OilPressure = data.OilPressure

	coordinate.Latitude = data.Latitude
	coordinate.Longitude = data.Longitude

	return vehicleData, coordinate

}
