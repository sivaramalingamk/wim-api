package domain

type VehicleData struct {
	ID               string `json:"id"`
	Weight           int    `json:"weight"`
	VehicleSpeed     int    `json:"vehicleSpeed"`
	Acceleration     int    `json:"acceleration"`
	HeadingDirection int    `json:"headingDirection"`
	CoolentTemp      int    `json:"coolentTemp"`
	OilPressure      int    `json:"oilPressure"`
	IntakeAirTemp    int    `json:"intake_airTemp"`
	Rpm              int    `json:"rpm"`
	EngineLoad       int    `json:"engineLoad"`
	ElevationAngle   int    `json:"elevationAngle"`
	O2               int    `json:"o2"`
	FuelFlow         int    `json:"fuelFlow"`
}
