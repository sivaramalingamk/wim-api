package domain

import "time"

type WeatherData struct {
	ID                 string    `json:"id"`
	Latitude           float64   `json:"latitude"`
	Longitude          float64   `json:"longitude"`
	Time               time.Time `json:"time"`
	WindSpeed          int       `json:"windSpeed"`
	WindDirection      int       `json:"windDirection"`
	AtmosphereTemp     int       `json:"atmosphereTemp"`
	AtmospherePressure int       `json:"atmospherePressure"`
	Humidity           int       `json:"humidity"`
}

type WeatherDataCollection struct {
	Wdc []WeatherData
}

func (wdc *WeatherDataCollection) AddData(data WeatherData) []WeatherData {
	wdc.Wdc = append(wdc.Wdc, data)
	return wdc.Wdc

}
