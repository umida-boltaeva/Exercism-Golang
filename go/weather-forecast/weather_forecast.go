// Package weather provides tools to fetch the current weather forecast.
package weather

// CurrentCondition represents current weather condition.
var CurrentCondition string

// CurrentLocation represents current location.
var CurrentLocation string

// Forecast returns a string that says the current weather condition at the given location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
