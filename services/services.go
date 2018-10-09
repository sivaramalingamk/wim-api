package services

import (
	"wim-api/domain"
	"wim-api/repository"
)

func ProcessVehicleData(data domain.VehicleData) (string, error) {

	if _, err := repository.AddVehicleData(data); err != nil {
		println("Error in Processing Vehicle Data")
		return "", err

	}

	return "Success", nil

}

func ProcessWetherData(data domain.WeatherData, id string) (string, error) {
	data.ID = id
	if _, err := repository.AddWeatherData(data); err != nil {
		println("Error in Processing Weather Data")
		return "", err

	}

	return "Success", nil
}

func ProcessTrainingData(data domain.TrainingData) (string, error) {

	if _, err := repository.AddTrainingData(data); err != nil {
		println("Error in Processing Training Data")
		return "", err

	}

	return "Success", nil

}
