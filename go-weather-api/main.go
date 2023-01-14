package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type apiConfig struct {
	OpenWeatherAPIKey string `json:"OpenWeatherAPIKey"`
}

type Weather struct {
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  float64 `json:"pressure"`
		Humidity  float64 `json:"humidity"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
		Deg   float64 `json:"deg"`
	} `json:"wind"`
	Name string `json:"name"`
}

func loadApiConfig(filename string) (apiConfig, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return apiConfig{}, err
	}

	var config apiConfig

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return apiConfig{}, err
	}

	return config, nil
}

func greetFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func queryWeather(city string) (weather, error) {
	config, err := loadApiConfig(".apiConfig.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + config.OpenWeatherAPIKey)
	if err != nil {
		return weather{}, err
	}

	defer resp.Body.Close()

	var weather Weather

	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		return weather{}, err
	}

	return weather, nil
}

func weatherFunc(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	if city == "" {
		http.Error(w, "Missing city parameter", http.StatusBadRequest)
		return
	}

	weather, err := queryWeather(city)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weather)
}

func main() {
	http.HandleFunc("/", greetFunc)
	http.HandleFunc("/weather", weatherFunc)
	http.ListenAndServe(":8080", nil)
}
