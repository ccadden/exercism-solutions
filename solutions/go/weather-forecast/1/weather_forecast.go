// Package weather provides the forcast in a given city.
package weather

// CurrentCondition is the weather conditions.
var CurrentCondition string

// CurrentLocation is the location of the forecast.
var CurrentLocation string

// Forecast returns a string of the current weather at a given location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
