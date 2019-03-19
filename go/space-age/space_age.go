// Given an age in seconds, calculate how old someone would be on:

// Earth: orbital period 365.25 Earth days, or 31557600 seconds
// Mercury: orbital period 0.2408467 Earth years
// Venus: orbital period 0.61519726 Earth years
// Mars: orbital period 1.8808158 Earth years
// Jupiter: orbital period 11.862615 Earth years
// Saturn: orbital period 29.447498 Earth years
// Uranus: orbital period 84.016846 Earth years
// Neptune: orbital period 164.79132 Earth years
// So if you were told someone were 1,000,000,000 seconds old, you should be able to say that they're 31.69 Earth-years old.

package space

// Planet type
type Planet string

// Age array
const (
	EarthSeconds float64 = 31557600   //seconds
	Earth        float64 = 365.25     // days
	Mercury      float64 = 0.2408467  // earth years
	Venus        float64 = 0.61519726 // earth years
	Mars         float64 = 1.8808158
	Jupiter      float64 = 11.862615
	Saturn       float64 = 29.447498
	Uranus       float64 = 84.016846
	Neptune      float64 = 164.7913
)

var names = [7]Planet{"Mercury", "Venus", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
var years = [7]float64{Mercury, Venus, Mars, Jupiter, Saturn, Uranus, Neptune}

// Age converts seconds to age of planet in earth years
func Age(seconds float64, planet Planet) float64 {
	var earthYear float64
	earthYear = seconds / EarthSeconds
	if planet == "Earth" {
		return earthYear
	}
	for i := 0; i < len(names); i++ {
		if planet == names[i] {
			return earthYear / years[i]
		}
	}
	return 0
}
