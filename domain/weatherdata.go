package domain

type WeatherData struct {
	ID                 string `json:"id"`
	WindSpeed          int    `json:"windSpeed"`
	WindDirection      int    `json:"windDirection"`
	AtmosphereTemp     int    `json:"atmosphereTemp"`
	AtmospherePressure int    `json:"atmospherePressure"`
	Humidity           int    `json:"humidity"`
}
