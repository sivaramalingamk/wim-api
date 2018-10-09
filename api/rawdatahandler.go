package api

import (
	"encoding/json"
	"github.com/go-chi/render"
	"math"
	"net/http"
	"wim-api/domain"
	"wim-api/io"
	"wim-api/services"
)

func SimpleDataHandler(w http.ResponseWriter, r *http.Request) {

	var data domain.RawInputData

	if err := render.Bind(r, &data); err != nil {
		w.Write([]byte("Error Binding Raw Data"))
		return
	}
	w.Write([]byte("Sending Vehicle data to Process ID :"))
	w.Write([]byte(data.ID))

	vdata, coord := bulkToSimple(data)

	w.Write([]byte("Vehicle data"))

	services.ProcessVehicleData(vdata)

	w.Write([]byte("Sending TO Weather API :"))
	owmRes, err1 := io.WeatherAPI(coord, vdata.ID)
	if err1 != nil {
		panic(err1)
	}
	services.ProcessWetherData(owmRes)
	w.Write([]byte(owmRes.ID))

	//services.ProcessWetherData(----)

	w.Write([]byte("Success at Simple vehicle data handler"))

}

func BulkVehicleDataHandler(w http.ResponseWriter, r *http.Request) {

	data := []domain.RawInputData{}
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		w.Write([]byte("Error on GETBULK DATA"))
		return
	}

	_, icoord := bulkToSimple(data[0])
	for i := range data {

		vdata, coord := bulkToSimple(data[i])
		services.ProcessVehicleData(vdata)
		if coordDiff(icoord, coord) { //to avoid frequent calls
			io.WeatherAPI(coord, vdata.ID)
			icoord = coord
		}

	}

	//services.ProcessVehicleData()

	//services.ProcessWetherData()

}

//@ if the coordinate difference is higher than the cordinate threshold then call weathermap api
func coordDiff(coordinate1 domain.Coordinate, coordinate2 domain.Coordinate) bool {

	if math.Abs(coordinate1.Latitude-coordinate2.Latitude) > io.Coord_Threshold || math.Abs(coordinate1.Longitude-coordinate2.Longitude) > io.Coord_Threshold {
		return true
	}

	return false
}
