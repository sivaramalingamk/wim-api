package domain

import "time"

type OutputData struct {
	ID           string    `json:"id"`
	Year         int       `json:"year"`
	Time         time.Time `json:"time"`
	StartLat     float64   `json:"startLat"`
	StartLon     float64   `json:"startLon"`
	EndLat       float64   `json:"EndLat"`
	EndLon       float64   `json:"EndLon"`
	Weight       int       `json:"weight"`
	IsOverloaded bool      `json:"isOverloaded"`
}

type OutputDataCollection struct {
	Odc []OutputData
}

func (odc *OutputDataCollection) AddData(data OutputData) []OutputData {
	odc.Odc = append(odc.Odc, data)
	return odc.Odc

}
