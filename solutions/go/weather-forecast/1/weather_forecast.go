// Package weather provides tools for weather.
package weather

var (
    // CurrentCondition represents the current condition of the weather.
	CurrentCondition string
    // CurrentLocation represents the location of target.
	CurrentLocation  string
)
// Forecast returns the forecast of the city's weather.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
