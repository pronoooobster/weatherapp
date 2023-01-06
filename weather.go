/*
	The application sends the request to openweathermap API and displays the weather in a specific location.
*/

package main

import (
	"bufio"
	"fmt"
	"os"

	owm "github.com/briandowns/openweathermap"
)

// ! change the API key to your own from openweathermap
var apiKey = "4a9b928d5ba52971890019bbc2b73967"

// display the weather right now
func printWeatherNow(weather *owm.CurrentWeatherData, location string) {
	// get the weather right now
	weather.CurrentByName(location)

	// check if the weather is not empty
	if len(weather.Weather) != 0 {
		// output the weather
		fmt.Println("Weather in", weather.Name, "|", weather.Sys.Country, "right now:")
		fmt.Println("Temperature: ", weather.Main.Temp, "°C", "| Feels like: ", weather.Main.FeelsLike, "°C")
		fmt.Println("Humidity: ", weather.Main.Humidity, "%")
		fmt.Println("Wind speed: ", weather.Wind.Speed, "m/s")
		fmt.Println("Weather description: ", weather.Weather[0].Description)
	} else {
		fmt.Println("The location is not found.")
	}
}

func main() {
	// create an API call
	weather, err := owm.NewCurrent("C", "en", apiKey)

	// handle the error
	if err != nil {
		fmt.Println("Error occured while making the API call")
		return
	}

	// get input from the user
	var location string
	fmt.Print("Enter the location: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		location = scanner.Text()
	}

	// display the weather right now
	printWeatherNow(weather, location)
}
