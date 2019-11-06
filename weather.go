package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"time"
)

type Weather struct {
	key string
}

type CurrentWeather struct {
	When         time.Time
	Location     string
	Temperature  int // degrees Celcius
	Descriptions []CurrentWeatherDescription
	Position     Position
}

type CurrentWeatherDescription struct {
	Icon string
	Text string
}

type owmCurrent struct {
	Cod   int `json:"cod"` // HTTP status code
	Coord struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"coord"`
	Main struct {
		Temp float64 `json:"temp"` // degrees Celcius
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"` // human-readable
		Icon        string `json:"icon"`        // OWM icon code
	} `json:"weather"`
	Name string `json:"name"` // location name (e.g. city)
}

func (w *Weather) Get(pos Position) (CurrentWeather, error) {
	weather, err := w.fetch(pos)

	if err != nil {
		return CurrentWeather{}, err
	}

	return weather, nil
}

func (w *Weather) fetch(pos Position) (CurrentWeather, error) {
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&units=metric&lang=de&appid=%s", pos.Lat, pos.Lon, os.Getenv("OPENWEATHERMAP_API_KEY"))

	log.Println(url)

	weather := owmCurrent{}

	if err := getJSON(url, &weather); err != nil {
		log.Printf("Failed fetching current weather from OpenWeatherMap API: %v", err)
		return CurrentWeather{}, err
	}

	if weather.Cod != 200 {
		log.Printf("Failed fetching current weather from OpenWeatherMap API: HTTP code %d", weather.Cod)
		return CurrentWeather{}, fmt.Errorf("HTTP code %d", weather.Cod)
	}

	cw := CurrentWeather{
		When:        time.Now(),
		Temperature: int(math.Round(weather.Main.Temp)),
		Location:    weather.Name,
		Position:    Position{weather.Coord.Lat, weather.Coord.Lon},
	}

	for _, description := range weather.Weather {
		cw.Descriptions = append(cw.Descriptions, CurrentWeatherDescription{
			Icon: description.Icon,
			Text: description.Description,
		})
	}

	return cw, nil
}
