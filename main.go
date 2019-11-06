package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type input struct {
	Positions []Position `json:"positions"`
}

type output struct {
	Results []result `json:"results"`
	Stats   stats    `json:"stats"`
	Errors  []string `json:"errors"`
}

type result struct {
	Input        Position `json:"input"`
	Output       Position `json:"output"`
	Distance     float64  `json:"distance_meters"` // meters
	LocationName string   `json:"location_name"`
}

type stats struct {
	AvgDistance float64 `json:"avg_distance_meters"` // meters
	MaxDistance float64 `json:"max_distance_meters"` // meters
	MinDistance float64 `json:"min_distance_meters"` // meters
}

func main() {
	if os.Getenv("OPENWEATHERMAP_API_KEY") == "" {
		log.Fatalf("Missing API key for https://openweathermap.org")
	}

	var i input

	err := json.NewDecoder(os.Stdin).Decode(&i)

	if err != nil {
		log.Fatalf("Failed to unmarshal JSON input: %v", err)
	}

	w := &Weather{}

	o := output{
		Results: []result{},
		Errors:  []string{},
	}

	var sumDistance float64

	for i, p := range i.Positions {
		if i > 0 {
			// Throttle API requests.
			// Free plan allows 60 requests per minute.
			// See: https://openweathermap.org/price
			time.Sleep(time.Second)
		}

		result, err := checkPosition(w, p)

		if err != nil {
			o.Errors = append(o.Errors, err.Error())
			continue
		}

		o.Results = append(o.Results, result)

		if result.Distance > o.Stats.MaxDistance {
			o.Stats.MaxDistance = result.Distance
		}

		if result.Distance < o.Stats.MinDistance || i == 0 {
			o.Stats.MinDistance = result.Distance
		}

		sumDistance += result.Distance
	}

	if len(o.Results) > 0 {
		o.Stats.AvgDistance = sumDistance / float64(len(o.Results))
	}

	j, err := json.MarshalIndent(o, "", "    ")

	if err != nil {
		log.Fatalf("Failed to marshal JSON output: %v", err)
	}

	fmt.Println(string(j))
}

func checkPosition(w *Weather, p Position) (result, error) {
	cw, err := w.Get(p)
	if err != nil {
		return result{}, err
	}

	return result{
		Input:        p,
		Output:       cw.Position,
		Distance:     p.DistanceTo(cw.Position),
		LocationName: cw.Location,
	}, nil
}
