package space

import (
	"math"
)

type Planet string

const (
	Mercury = "Mercury"
	Venus   = "Venus"
	Earth   = "Earth"
	Mars    = "Mars"
	Jupiter = "Jupiter"
	Saturn  = "Saturn"
	Uranus  = "Uranus"
	Neptune = "Neptune"
)

var planets = map[Planet]float64{
	Earth: 31557600,
}

func init() {
	planets[Mercury] = planets[Earth] * 0.2408467
	planets[Venus] = planets[Earth] * 0.61519726
	planets[Mars] = planets[Earth] * 1.8808158
	planets[Jupiter] = planets[Earth] * 11.862615
	planets[Saturn] = planets[Earth] * 29.447498
	planets[Uranus] = planets[Earth] * 84.016846
	planets[Neptune] = planets[Earth] * 164.79132
}

func Age(seconds float64, planet Planet) float64 {
	if ratio, ok := planets[planet]; ok {
		return (math.Round((seconds/ratio)*100) / 100)
	}
	panic("undefined value " + planet + " with no known year ratio")
}
