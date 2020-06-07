package space

type Planet string

const EarthSecondsPerYear = 31557600

func Age(seconds float64, planet Planet) float64 {
	if planet == "Mercury" {
		return seconds / 0.2408467 / EarthSecondsPerYear
	}

	if planet == "Venus" {
		return seconds / 0.61519726 / EarthSecondsPerYear
	}

	if planet == "Earth" {
		return seconds / 1 / EarthSecondsPerYear
	}

	if planet == "Mars" {
		return seconds / 1.8808158 / EarthSecondsPerYear
	}

	if planet == "Jupiter" {
		return seconds / 11.862615 / EarthSecondsPerYear
	}

	if planet == "Saturn" {
		return seconds / 29.447498 / EarthSecondsPerYear
	}

	if planet == "Uranus" {
		return seconds / 84.016846 / EarthSecondsPerYear
	}

	if planet == "Neptune" {
		return seconds / 164.79132 / EarthSecondsPerYear
	}

	return 0
}
