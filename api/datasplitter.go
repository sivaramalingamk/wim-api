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

func TrainDataMerger(vdata domain.VehicleData, wdata domain.WeatherData) domain.TrainingData {
	var tdata domain.TrainingData

	tdata.ID = vdata.ID
	tdata.AtmospherePressure = wdata.AtmospherePressure
	tdata.WindSpeed = wdata.WindSpeed
	tdata.WindDirection = wdata.WindDirection
	tdata.Humidity = wdata.Humidity
	tdata.AtmosphereTemp = wdata.AtmosphereTemp
	tdata.FuelFlow = vdata.FuelFlow
	tdata.O2 = vdata.O2
	tdata.ElevationAngle = vdata.ElevationAngle
	tdata.EngineLoad = vdata.EngineLoad
	tdata.Rpm = vdata.Rpm
	tdata.IntakeAirTemp = vdata.IntakeAirTemp
	tdata.OilPressure = vdata.OilPressure
	tdata.CoolantTemp = vdata.CoolentTemp
	tdata.HeadingDirection = vdata.HeadingDirection
	tdata.Acceleration = vdata.Acceleration
	tdata.Weight = vdata.Weight
	tdata.VehicleSpeed = vdata.VehicleSpeed
	return tdata
}
