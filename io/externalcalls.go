package io

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"wim-api/domain"
)

func WeatherAPI(coordinate domain.Coordinate) string {
	baseUrl := "https://api.openweathermap.org/data/2.5/weather"
	appid := API_Key
	url1 := fmt.Sprintf(baseUrl+"?lat=%d&lon=%d&appid=%s", coordinate.Latitude, coordinate.Longitude, appid)
	fmt.Println("URL:>", url1)

	req, err := http.NewRequest("GET", url1, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	//fmt.Println("REquest=",req)
	return string(body)
}
