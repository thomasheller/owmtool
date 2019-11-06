package main

import "math"

const earthRadius float64 = 6371000 // meters

type Position struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

func (p Position) toRadian() Position {
	return Position{
		Lat: p.toRadians(p.Lat),
		Lon: p.toRadians(p.Lon),
	}
}

func (p Position) toRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func (p Position) delta(other Position) Position {
	return Position{
		Lat: p.Lat - other.Lat,
		Lon: p.Lon - other.Lon,
	}
}

func (p Position) DistanceTo(other Position) float64 {
	here := p.toRadian()
	there := other.toRadian()

	change := here.delta(there)

	a := math.Pow(math.Sin(change.Lat/2), 2) + math.Cos(here.Lat)*math.Cos(there.Lat)*math.Pow(math.Sin(change.Lon/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return c * earthRadius
}
