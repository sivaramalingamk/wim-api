package repository

import (
	"fmt"
	"time"
	"wim-api/domain"
)

func AddHistoricalWeatherData(data domain.WeatherData) (string, error) {
	fmt.Println(" **** Creating new Historical Weather data ****\n", data)
	defer getCluster().Close()
	if err := Session.Query("INSERT INTO historicweatherdata(time,atmospherepressure,atmospheretemp,humidity,winddirection,windspeed) values(?,?,?,?,?,?)", time.Now(), data.AtmospherePressure, data.AtmosphereTemp, data.Humidity, data.WindDirection, data.WindSpeed).Exec(); err != nil {

		fmt.Println("Error while inserting Weather Data", err)
		fmt.Println(err)
		return "", err
	}

	return "Success", nil
}

/*
func SelectHistoricalWeatherData() (domain.WeatherDataCollection, string) {
	fmt.Println("Selecting Weather Data colection From Database")
	defer getCluster().Close()
	wDataCollection := domain.WeatherDataCollection{}
	temp := domain.WeatherData{}
	iter := Session.Query(`SELECT time,atmospherepressure,atmospheretemp,humidity,winddirection,windspeed FROM weatherdata`).Iter()
	for iter.Scan(&temp.ID, &temp.AtmospherePressure, &temp.AtmosphereTemp, &temp.Humidity, &temp.WindDirection, &temp.WindSpeed, &temp.Latitude, &temp.Longitude) {
		wDataCollection.AddData(temp)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
		return wDataCollection, "Error"
	}
	fmt.Println("Selected ", len(wDataCollection.Wdc), " Number of Weather Data")
	return wDataCollection, "Success"
}
*/
