package io

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"wim-api/domain"
)

func WeatherAPI(coordinate domain.Coordinate) (string, error) {
	baseUrl := "https://api.openweathermap.org/data/2.5/weather"
	appid := API_Key
	url1 := fmt.Sprintf(baseUrl+"?lat=%f&lon=%f&appid=%s", coordinate.Latitude, coordinate.Longitude, appid)
	fmt.Println("URL:>", url1)

	req, err := http.NewRequest("GET", url1, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	var wdata map[string]interface{}
	json.Unmarshal(body, &wdata)

	//fmt.Println("REquest=",req)
	return string(body), err
}
