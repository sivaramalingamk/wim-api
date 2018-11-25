package api

import (
	"fmt"
	"github.com/robfig/cron"
	"wim-api/domain"
	"wim-api/io"
	"wim-api/repository"
	"wim-api/services"
)

func MergeAndInsertTraining() {
	//	data:=[]domain.TrainingData{}
	c := cron.New()
	c.AddFunc("@every 5m", func() {
		fmt.Println("Testing Cron")
		tdc, _ := merger()
		services.ProcessTrainingData(tdc)

	})
	c.Start()

}

func PopulateHistoricWeather() {
	coordinate := domain.Coordinate{}
	coordinate.Latitude = -33.918861
	coordinate.Longitude = 18.423300

	c := cron.New()
	c.AddFunc("@every 10m", func() {
		fmt.Println("Getting Current Weather Data for Cape Town")
		owmData, err := io.WeatherAPI(coordinate)
		if err != nil {
			fmt.Println("Error :", err, owmData.ID)

		} else {
			fmt.Println("Writing into database")
			services.ProcessHistoricalWeatherData(owmData)
		}

	})
	c.Start()
}
func merger() (domain.TrainingDataCollection, string) {
	tdc := domain.TrainingDataCollection{}

	wdc := make(chan domain.WeatherDataCollection)
	vdc := make(chan domain.VehicleDataCollection)

	go func() { wdc <- getWeatherData() }()
	go func() { vdc <- getVehicleData() }()

	wdC := <-wdc
	vdC := <-vdc

	/*
	 vdC,_:=repository.SelectAllVehicle()
	 wdC,_:=repository.SelectWeatherData()
	*/
	fmt.Print("Lenght of WDC=", len(wdC.Wdc))
	fmt.Print("Lenght of vDC=", len(vdC.Vdc))
	for _, wd := range wdC.Wdc {
		for _, vd := range vdC.Vdc {
			if vd.ID == wd.ID {
				fmt.Println("Adding values VehicleID:", vd.ID, "WeatherID:", wd.ID)
				tdc.AddData(trainDataMerger(vd, wd))
			}
		}

	}

	return tdc, "Success"

}

func getWeatherData() domain.WeatherDataCollection {

	wdc, _ := repository.SelectWeatherData()

	return wdc
}

func getVehicleData() domain.VehicleDataCollection {

	vdc, _ := repository.SelectAllVehicle()
	return vdc
}

func trainDataMerger(vdata domain.VehicleData, wdata domain.WeatherData) domain.TrainingData {
	var tdata domain.TrainingData

	tdata.ID = vdata.ID
	tdata.AtmospherePressure = wdata.AtmospherePressure
	tdata.WindSpeed = wdata.WindSpeed
	tdata.WindDirection = wdata.WindDirection
	tdata.Humidity = wdata.Humidity
	tdata.AtmosphereTemp = wdata.AtmosphereTemp
	tdata.FuelFlow = vdata.FuelFlow
	tdata.O2 = vdata.O2
	tdata.ElevationAngle = vdata.ElevationAngle
	tdata.EngineLoad = vdata.EngineLoad
	tdata.Rpm = vdata.Rpm
	tdata.IntakeAirTemp = vdata.IntakeAirTemp
	tdata.OilPressure = vdata.OilPressure
	tdata.CoolantTemp = vdata.CoolentTemp
	tdata.HeadingDirection = vdata.HeadingDirection
	tdata.Acceleration = vdata.Acceleration
	tdata.Weight = vdata.Weight
	tdata.VehicleSpeed = vdata.VehicleSpeed
	return tdata
}
