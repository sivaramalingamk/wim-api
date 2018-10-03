package domain

type RawInputData struct {
	ID               string `json:"id"`
	Time             string `json:"time"`
	Latitude         string `json:"latitude"`
	Longitude        string `json:"longitude"`
	Weight           int    `json:"weight"`
	VehicleSpeed     int    `json:"vehicleSpeed"`
	Acceleration     int    `json:"acceleration"`
	HeadingDirection int    `json:"headingDirection"`
	CoolentTemp      int    `json:"coolentTemp"`
	OilPressure      int    `json:"oilPressure"`
	IntakeAirTemp    int    `json:"intakeAirTemp"`
	Rpm              int    `json:"rpm"`
	EngineLoad       int    `json:"engineLoad"`
	ElevationAngle   int    `json:"elevationAngle"`
	O2               int    `json:"o2"`
	FuelFlow         int    `json:"fuelFlow"`
}
