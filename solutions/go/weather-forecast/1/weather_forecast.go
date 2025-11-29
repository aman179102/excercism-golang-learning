// Package weather provides tools to get the current weather
// condition and location for cities in Goblinocus.
package weather

var (
    // CurrentCondition stores the current weather condition for the selected city.
    CurrentCondition string

    // CurrentLocation stores the name of the city for which the weather is being reported.
    CurrentLocation string
)

// Forecast returns a string describing the current weather condition
// for the given city in Goblinocus.
func Forecast(city, condition string) string {
    CurrentLocation, CurrentCondition = city, condition
    return CurrentLocation + " - current weather condition: " + CurrentCondition
}
