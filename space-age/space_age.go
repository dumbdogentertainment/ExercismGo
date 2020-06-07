package space

type Planet string

func Age(seconds float64, planet Planet) float64 {
	if planet == "Mercury" {
		return seconds / 0.2408467 / 31557600
	}

	if planet == "Venus" {
		return seconds / 0.61519726 / 31557600
	}

	if planet == "Earth" {
		return seconds / 1 / 31557600
	}

	if planet == "Mars" {
		return seconds / 1.8808158 / 31557600
	}

	if planet == "Jupiter" {
		return seconds / 11.862615 / 31557600
	}

	if planet == "Saturn" {
		return seconds / 29.447498 / 31557600
	}

	if planet == "Uranus" {
		return seconds / 84.016846 / 31557600
	}

	if planet == "Neptune" {
		return seconds / 164.79132 / 31557600
	}

	return 0
}
