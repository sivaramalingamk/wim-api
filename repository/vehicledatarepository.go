package repository

import (
	"fmt"
	"wim-api/domain"
)

func AddVehicleData(data domain.VehicleData) (string, error) {
	fmt.Println(" **** Creating new data ****\n", data)
	defer getCluster().Close()
	if err := Session.Query("INSERT INTO vehicledata(id, weight,vehiclespeed,acceleration,headingdirection,coolenttemp,oilpressure,intakeairtemp,rpm,engineload,elevationangle,o2,fuelflow) values(?,?,?,?,?,?,?,?,?,?,?,?,?)",
		data.ID, data.Weight, data.VehicleSpeed, data.Acceleration, data.HeadingDirection, data.CoolentTemp, data.OilPressure, data.OilPressure, data.IntakeAirTemp, data.Rpm, data.EngineLoad, data.ElevationAngle, data.O2, data.FuelFlow).Exec(); err != nil {

		fmt.Println("Error while inserting Vehicle Data", err)
		fmt.Println(err)
		return "", err
	}

	return "Success", nil
}
