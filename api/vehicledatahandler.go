package api

import (
	"github.com/go-chi/render"
	"net/http"
	"wim-api/domain"
	"wim-api/services"
)

func VehicleDataHandler(w http.ResponseWriter, r *http.Request) {

	var data domain.RawInputData

	if err := render.Bind(r, &data); err != nil {
		w.Write([]byte("Error Binding Raw Data"))
		return
	}
	w.Write([]byte("Sending Vehicle data to Process ID :"))
	w.Write([]byte(data.ID))

	var vehicleData domain.VehicleData

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

	w.Write([]byte("Vehicle data"))

	services.ProcessVehicleData(vehicleData)

	// services.ProcessWetherData(weatherData)

	w.Write([]byte("Success at Simple vehicle data handler"))

}

func BulkVehicleDataHandler(w http.ResponseWriter, r *http.Request) {

	var data []domain.RawInputData

	if err := render.Bind(r, &data); err != nil {
		w.Write([]byte("Error Binding Raw Data"))
		return
	}
	w.Write([]byte("Sending Vehicle data to Process ID :"))
	w.Write([]byte(data.ID))
	var vehicleData domain.VehicleData

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

	w.Write([]byte("Vehicle data"))

	services.ProcessVehicleData(vehicleData)

	// services.ProcessWetherData(weatherData)

	w.Write([]byte("Success at Simple vehicle data handler"))

	//services.ProcessVehicleData()

	//services.ProcessWetherData()
	w.Write([]byte("Success at Simple vehicle data handler"))
}
