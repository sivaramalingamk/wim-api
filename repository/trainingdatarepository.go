package repository

import (
	"fmt"
	"wim-api/domain"
)

func AddTrainingData(data domain.TrainingData) (string, error) {
	fmt.Println(" **** Creating new data ****\n", data)
	defer getCluster().Close()
	if err := Session.Query("INSERT INTO trainingdata(id, weight,vehiclespeed,headingdirection,winddirection,atmospheretemp,atmospherepressure,coolanttemp,oilpressure,intakeairtemp,humidity,rpm,engineload,elevationangle,o2,fuelflow) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		data.ID, data.Weight, data.VehicleSpeed, data.Acceleration, data.WindSpeed, data.HeadingDirection, data.WindDirection, data.AtmosphereTemp, data.AtmospherePressure, data.CoolantTemp, data.OilPressure, data.IntakeAirTemp, data.Humidity, data.Rpm, data.EngineLoad, data.ElevationAngle, data.O2, data.FuelFlow).Exec(); err != nil {

		fmt.Println("Error while inserting Vehicle Data", err)
		fmt.Println(err)
		return "", err
	}

	return "Success", nil
}
