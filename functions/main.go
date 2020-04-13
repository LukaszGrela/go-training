package main

import "fmt"

func kelvinToCelcius(k float64) float64 {
	return k - 273.15
}
func celciusToFahrenheit(c float64) float64 {
	return (c * 9.0 / 5.0) + 32.0
}
func kelvinToFahrenheit(k float64) float64 {
	return celciusToFahrenheit(kelvinToCelcius(k))
}

func main() {
	// Keyword Func
	//         Name  Parameter  result
	//  func   Intn   (n  int)    int
	//               name type   type
	//

	kelvin := 294.0
	celcius := kelvinToCelcius(kelvin)
	fmt.Println(kelvin, "°K is ", celcius, "°C")

	kelvin = 233.0
	fahrenheit := kelvinToFahrenheit(kelvin)
	fmt.Println(kelvin, "°K is ", fahrenheit, "°F")
	fmt.Println(0, "°K is ", kelvinToFahrenheit(0), "°F")
}
