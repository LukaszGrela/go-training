package main

import "fmt"

type celcius float64
type fahrenheit float64

/// celcius
// fahrenheit converts °C to °F
func (c celcius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

// celcius converts °F to °C
func (f fahrenheit) celcius() celcius {
	return celcius((f - 32.0) * 5.0 / 9.0)
}

type getRowFn func(input float64) (string, string)

func formatRow(a, b string) {
	fmt.Printf("| %6s     | %6s     |\n", a, b)
}

func celciusToFahrenheitRow(input float64) (string, string) {

	c := celcius(input)
	f := c.fahrenheit()

	return fmt.Sprintf("%.1f", c), fmt.Sprintf("%.1f", f)
}
func fahrenheitToCelciusRow(input float64) (string, string) {

	f := fahrenheit(input)
	c := f.celcius()

	return fmt.Sprintf("%.1f", f), fmt.Sprintf("%.1f", c)
}

func drawTable(h1, h2 string, start float64, step float64, count float64, f getRowFn) {
	fmt.Printf("=============+=============\n")
	formatRow(h1, h2)
	fmt.Printf("=============+=============\n")
	for i := start; i <= count; i += step {
		col1, col2 := f(i)
		formatRow(col1, col2)
	}
	fmt.Printf("=============+=============\n")
}

func celciusToFahrenheitTable() {

	drawTable("°C", "°F", -40.0, 5.0, 100.0, celciusToFahrenheitRow)

}

func fahrenheitToCelciusTable() {

	drawTable("°F", "°C", -40.0, 5.0, 100.0, fahrenheitToCelciusRow)

}

func main() {
	celciusToFahrenheitTable()
	fmt.Println()
	fahrenheitToCelciusTable()
}
