package repository

import (
	"fmt"
	"wim-api/domain"
)

func AddWeatherData(data domain.WeatherData) (string, error) {
	fmt.Println(" **** Creating new data ****\n", data)
	defer getCluster().Close()
	if err := Session.Query("INSERT INTO weatherdata(id,atmospherepressure,atmospheretemp,humidity,winddirection,windspeed) values(?,?,?,?,?,?)", data.ID, data.AtmospherePressure, data.AtmosphereTemp, data.Humidity, data.WindDirection, data.WindSpeed).Exec(); err != nil {

		fmt.Println("Error while inserting Weather Data", err)
		fmt.Println(err)
		return "", err
	}

	return "Success", nil
}
