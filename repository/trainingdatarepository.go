package repository

import (
	"fmt"
	"wim-api/domain"
)

func AddTrainingData(td domain.TrainingData) (string, error) {
	fmt.Println(" **** Creating new data ****\n", td)
	defer getCluster().Close()
	if err := Session.Query("INSERT INTO trainingdata( id,acceleration, atmospherepressure, atmospheretemp, coolanttemp, elevationangle, engineload, fuelflow, headingdirection, humidity, intakeairtemp, o2, oilpressure, rpm, vehiclespeed, weight, winddirection, windspeed ) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		td.ID, td.Acceleration, td.AtmospherePressure, td.AtmosphereTemp, td.CoolantTemp, td.ElevationAngle, td.EngineLoad, td.FuelFlow, td.HeadingDirection, td.Humidity, td.IntakeAirTemp, td.O2, td.OilPressure, td.Rpm, td.VehicleSpeed, td.Weight, td.WindDirection, td.WindSpeed).Exec(); err != nil {

		fmt.Println("Error while inserting Training Data", err)
		fmt.Println(err)
		return "", err
	}

	return "Success", nil
}
