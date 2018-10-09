package repository

import (
	"fmt"
	"wim-api/domain"
)

func AddWeatherData(data domain.WeatherData) (string, error) {
	fmt.Println(" **** Creating new data ****\n", data)
	defer getCluster().Close()
	if err := Session.Query("INSERT INTO weatherdata(id,atmospheretemp,humidity,windspeed,atmospherepressure,winddirection) values(?,?,?,?,?)",
		data.ID, data.AtmosphereTemp, data.Humidity, data.WindSpeed,
		data.AtmospherePressure, data.WindDirection).Exec(); err != nil {

		fmt.Println("Error while inserting Vehicle Data", err)
		fmt.Println(err)
		return "", err
	}

	return "Success", nil
}
