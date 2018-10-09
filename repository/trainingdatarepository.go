package repository

import (
	"fmt"
	"log"
	"wim-api/api"
	"wim-api/domain"
)

func AddTrainingData(data domain.TrainingData) (string, error) {
	fmt.Println(" **** Creating new data ****\n", data)
	defer getCluster().Close()
	if err := Session.Query("INSERT INTO trainingdata(id, weight,vehiclespeed,headingdirection,winddirection,atmospheretemp,atmospherepressure,coolanttemp,oilpressure,intakeairtemp,humidity,rpm,engineload,elevationangle,o2,fuelflow) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		data.ID, data.Weight, data.VehicleSpeed, data.Acceleration, data.WindSpeed, data.HeadingDirection, data.WindDirection, data.AtmosphereTemp, data.AtmospherePressure, data.CoolantTemp, data.OilPressure, data.IntakeAirTemp, data.Humidity, data.Rpm, data.EngineLoad, data.ElevationAngle, data.O2, data.FuelFlow).Exec(); err != nil {

		fmt.Println("Error while inserting Training Data", err)
		fmt.Println(err)
		return "", err
	}

	return "Success", nil
}

func SelectMerging() []domain.TrainingData {
	defer getCluster().Close()
	wdata := domain.WeatherData{}
	vdata := domain.VehicleData{}
	tdata := []domain.TrainingData{}
	var i = 0
	iter := Session.Query(`SELECT id,atmospherepressure,atmospheretemp,humidity,winddirection,windspeed FROM weatherdata`).Iter()
	for iter.Scan(wdata) {
		fmt.Println("Data ID:%s , atmP:%d ", wdata.ID, wdata.AtmospherePressure)
		iter2 := Session.Query("Select * from vehicledata where id=%s", wdata.ID)
		iter2.Scan(vdata)
		tdata[i] = api.TrainDataMerger(vdata, wdata)
		i++
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
	return tdata
}
