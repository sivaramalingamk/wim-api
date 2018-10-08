package api

import (
	"encoding/json"
	"github.com/go-chi/render"
	"math"
	"net/http"
	"strconv"
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
	w.Write([]byte(io.WeatherAPI(coord)))

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
		if coordDiff(icoord, coord) {
			io.WeatherAPI(coord)
			icoord = coord
		}

	}

	//services.ProcessVehicleData()

	//services.ProcessWetherData()

}

func coordDiff(coordinate1 domain.Coordinate, coordinate2 domain.Coordinate) bool {

	la1, e1 := stf(coordinate1.Latitude)
	la2, e2 := stf(coordinate2.Latitude)

	lo1, e3 := stf(coordinate1.Longitude)
	lo2, e4 := stf(coordinate2.Longitude)

	if e1 == nil && e2 == nil && e3 == nil && e4 == nil {
		if math.Abs(la1-la2) > 5 || math.Abs(lo1-lo2) > 5 {
			return true
		}
	}
	return false
}

func stf(string2 string) (float64, error) {
	val, err := strconv.ParseFloat(string2, 64)
	if err != nil {
		return 0, err
	}

	return val, nil
}
