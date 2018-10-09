package services

import (
	"wim-api/domain"
	"wim-api/repository"
)

func ProcessVehicleData(data domain.VehicleData) (string, error) {
	var err error
	if _, err := repository.AddVehicleData(data); err != nil {
		println("Error in Processing Vehicle Data")
		return "", nil

	}

	return "Success", err

}

func ProcessWetherData(data domain.WeatherData) (string, error) {
	var err error
	if _, err := repository.AddWeatherData(data); err != nil {
		println("Error in Processing Training Data")
		return "", nil

	}

	return "Success", err
}

func ProcessTrainingData(data domain.TrainingData) (string, error) {

	var err error
	if _, err := repository.AddTrainingData(data); err != nil {
		println("Error in Processing Training Data")
		return "", nil

	}

	return "Success", err

}
