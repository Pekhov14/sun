package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Weather struct {
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp     float64 `json:"temp"`
		Pressure int     `json:"pressure"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
	City string `json:"name"`
}

func main() {
	responce, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=Kiev&appid=8c37253bf3ce85a0b2bd9ed9b8b82e3e&units=metric")
	if err != nil {
		panic(err)
	}

	defer responce.Body.Close()

	if responce.StatusCode != 200 {
		panic("Weather API not avaliabe")
	}

	body, err := io.ReadAll(responce.Body)
	if err != nil {
		panic(err)
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}

	city, weatherCondition, temperatyre, wind := weather.City, weather.Weather[0], weather.Main, weather.Wind

	fmt.Printf(
		"%s, %.2fC, weatherCondition (%s, %s).\n",
		city,
		temperatyre.Temp,
		weatherCondition.Main,
		weatherCondition.Description,
	)

	fmt.Printf(
		"Humidity - %d, Pressure - %d, wind speed 🌬 - %.2f",
		temperatyre.Humidity,
		temperatyre.Pressure,
		wind.Speed,
	)
}
