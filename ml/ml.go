package ml

import (
	"fmt"
	"github.com/sivaramalingamk/regression"
	"wim-api/domain"
)

func Regression(collection domain.TrainingDataCollection) {

	r := new(regression.Regression)

	r.SetObserved("Weight")
	r.SetVar(0, "Acceleration")
	r.SetVar(1, "Vehicle_Speed")
	r.SetVar(2, "Heading_Direction")
	r.SetVar(3, "Elevation_Angle")
	r.SetVar(4, "Engine_Load")
	r.SetVar(5, "Coolent_Temp")
	r.SetVar(6, "Oil_Pressure")

	r.SetVar(7, "O2_Flow")
	r.SetVar(8, "Fuel_Flow")

	r.SetVar(9, "RPM")
	r.SetVar(10, "Intake_Air_Temp")
	r.SetVar(11, "Atmosphere_Temp")
	//r.SetVar(12, "Atmosphere_Pressure")
	//r.SetVar(13, "Humidity")
	//r.SetVar(14, "Wind_Speed")
	//r.SetVar(15, "Wind_Direction")

	for _, td := range collection.Tdc {
		r.Train(regression.DataPoint(float64(td.Weight), []float64{float64(td.Acceleration), float64(td.VehicleSpeed), float64(td.HeadingDirection), float64(td.ElevationAngle), float64(td.EngineLoad), float64(td.CoolantTemp), float64(td.OilPressure), float64(td.O2), float64(td.FuelFlow), float64(td.Rpm), float64(td.IntakeAirTemp), float64(td.AtmosphereTemp) /*,float64(td.AtmospherePressure),float64(td.Humidity),float64(td.WindSpeed),float64(td.WindDirection)*/}))
	}

	/*
		r.SetObserved("Murders per annum per 1,000,000 inhabitants")
		r.SetVar(0, "Inhabitants")
		r.SetVar(1, "Percent with incomes below $5000")
		r.SetVar(2, "Percent unemployed")
		r.Train(
				regression.DataPoint(11.2, []float64{587000, 16.5, 6.2}))
		r.Train(
				regression.DataPoint(13.4, []float64{643000, 20.5, 6.4}))
		r.Train(
				regression.DataPoint(40.7, []float64{635000, 26.3, 9.3}))
		r.Train(
				regression.DataPoint(5.3, []float64{692000, 16.5, 5.3}))
		r.Train(
				regression.DataPoint(24.8, []float64{1248000, 19.2, 7.3}))
		r.Train(
				regression.DataPoint(12.7, []float64{643000, 16.5, 5.9}))
		r.Train(
				regression.DataPoint(20.9, []float64{1964000, 20.2, 6.4}))
		r.Train(
				regression.DataPoint(35.7, []float64{1531000, 21.3, 7.6}))
		r.Train(
				regression.DataPoint(8.7, []float64{713000, 17.2, 4.9}))
		r.Train(
				regression.DataPoint(9.6, []float64{749000, 14.3, 6.4}))
		r.Train(
				regression.DataPoint(14.5, []float64{7895000, 18.1, 6}))
		r.Train(
				regression.DataPoint(26.9, []float64{762000, 23.1, 7.4}))
		r.Train(
				regression.DataPoint(15.7, []float64{2793000, 19.1, 5.8}))
		r.Train(
				regression.DataPoint(36.2, []float64{741000, 24.7, 8.6}))
		r.Train(
				regression.DataPoint(18.1, []float64{625000, 18.6, 6.5}))
		r.Train(
				regression.DataPoint(28.9, []float64{854000, 24.9, 8.3}))
		r.Train(
				regression.DataPoint(14.9, []float64{716000, 17.9, 6.7}))
		r.Train(
				regression.DataPoint(25.8, []float64{921000, 22.4, 8.6}))
		r.Train(
				regression.DataPoint(21.7, []float64{595000, 20.2, 8.4}))
		r.Train(
				regression.DataPoint(25.7, []float64{3353000, 16.9, 6.7}))
	*/
	r.AddCross(regression.MultiplierCross(1, 2))
	r.AddCross(regression.MultiplierCross(1, 3))
	r.AddCross(regression.MultiplierCross(2, 3, 4))
	r.AddCross(regression.MultiplierCross(3, 4))
	r.Run()

	fmt.Printf("Regression formula:\n%v\n", r.Formula)
	fmt.Printf("Regression:\n%s\n", r)

	for i, c := range r.Coef {
		if i == 0 {

			continue
		}
		if c < 0 {
			fmt.Println("Coefficient is negative, but shouldn't be: %.2f", c)
		}
	}

	if r.R2 < 0.8 {
		fmt.Println("R^2 was %.2f, but we expected > 80", r.R2)
	}

	prediction, err := r.Predict([]float64{1.01, 15.50, 30.00, 19.43, 4.39, 50.00, 1220.00, 34.00, 24.00, 999.00, 233.00, 303.20})
	fmt.Println("Predicted weight is :", prediction, " With error: ", err)
}
