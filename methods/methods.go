package main

import "fmt"

type celcius float64
type fahrenheit float64
type kelvin float64

// ceclius converts °K to °C
func (k kelvin) celcius() celcius {
	return celcius(k - 273.15)
}

// fahrenheit converts °K to °F
func (k kelvin) fahrenheit() fahrenheit {
	return fahrenheit(((k - 273.15) * 9.0 / 5.0) + 32.0)
}

// celcius converts °F to °C
func (f fahrenheit) celcius() celcius {
	return celcius((f - 32.0) * 5.0 / 9.0)
}

// kelvin converts °F to °K
func (f fahrenheit) kelvin() kelvin {
	return kelvin((f + 459.67) * 5.0 / 9.0)
}

// fahrenheit converts °C to °F
func (c celcius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

// kelvin converts °C to °K
func (c celcius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

func kelvinToCelcius(k kelvin) celcius {
	return celcius(k - 273.15)
}
func celciusToFahrenheit(c celcius) fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}
func kelvinToFahrenheit(k kelvin) fahrenheit {
	return celciusToFahrenheit(kelvinToCelcius(k))
}
func celciusToKelvin(c celcius) kelvin {
	return kelvin(c + 273.15)
}

func main() {
	//                     Method
	// Keyword  receiver    Name      result
	//  func   (k kelvin) celcius()  celcius
	//        name type                type
	//
	var temperature celcius = 20

	fmt.Println(temperature)

	var k kelvin = 294.0
	c := kelvinToCelcius(k)
	fmt.Println(k, "°K is ", c, "°C")
	fmt.Println(k, "°K is ", k.celcius(), "°C")
	fmt.Println(k, "°K is ", k.fahrenheit(), "°F")
	fmt.Println(k, "°K is ", kelvinToFahrenheit(k), "°F")

	var sunlitMoonSurface celcius = 127

	fmt.Println(sunlitMoonSurface, "°C of sunlit moon surface is ", celciusToKelvin(sunlitMoonSurface), "°K")

	fmt.Println(303.15, "°K is ", kelvin(303.15).celcius(), "°C")
	fmt.Println(303.15, "°K is ", kelvin(303.15).fahrenheit(), "°F")
	fmt.Println(86, "°F is ", fahrenheit(86).celcius(), "°C")
	fmt.Println(86, "°F is ", fahrenheit(86).kelvin(), "°K")
	fmt.Println(30, "°C is ", celcius(30).fahrenheit(), "°F")
	fmt.Println(30, "°C is ", celcius(30).kelvin(), "°K")

}
