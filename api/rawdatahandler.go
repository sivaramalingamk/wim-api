package api

import (
	"encoding/json"
	"fmt"
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

	msg, err2 := services.ProcessVehicleData(vdata)
	if err2 != nil {
		fmt.Print("Error occured :", err2)
	}
	fmt.Println("message from svc.procesvehicledata: ", msg)

	w.Write([]byte("\nSending TO Weather API :"))
	owmRes, err1 := io.WeatherAPI(coord)
	if err1 != nil {
		panic(err1)
	}
	services.ProcessWetherData(owmRes, vdata.ID)
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
	wdata, err := io.WeatherAPI(icoord)
	if err != nil {
		w.Write([]byte("Error on Weather API"))
	}

	for i := range data {
		vdata, coord := bulkToSimple(data[i])
		services.ProcessVehicleData(vdata)

		if coordDiff(icoord, coord) { //to avoid frequent calls
			wdata1, err := io.WeatherAPI(coord)
			if err != nil {
				w.Write([]byte("Error on Weather API"))
			}
			services.ProcessWetherData(wdata1, data[i].ID)
			icoord = coord
			wdata = wdata1
		} else {
			services.ProcessWetherData(wdata, data[i].ID)
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
