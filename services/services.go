package services

import (
	"wim-api/domain"
	"wim-api/repository"
)

func ProcessVehicleData(data domain.VehicleData) (string, error) {
	var err error
	if _, err := repository.AddVehicleData(data); err != nil {
		panic("Error in Processing Vehicle Data")
		return "", nil

	}

	return "Success", err

}

func ProcessWetherData(data domain.WeatherData) {

}

func ProcessTrainingData(data domain.TrainingData) error {
	return nil
}
