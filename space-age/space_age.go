package space

type Planet string

const EarthSecondsPerYear = 31557600

var PlanetarySecondsPerYear = map[Planet]float64{
	"Mercury": 0.2408467 * EarthSecondsPerYear,
	"Venus":   0.61519726 * EarthSecondsPerYear,
	"Earth":   EarthSecondsPerYear,
	"Mars":    1.8808158 * EarthSecondsPerYear,
	"Jupiter": 11.862615 * EarthSecondsPerYear,
	"Saturn":  29.447498 * EarthSecondsPerYear,
	"Uranus":  84.016846 * EarthSecondsPerYear,
	"Neptune": 164.79132 * EarthSecondsPerYear,
}

func Age(seconds float64, planet Planet) float64 {
	if val, exists := PlanetarySecondsPerYear[planet]; exists {
		return seconds / val
	}

	return 0
}
