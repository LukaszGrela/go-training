package main

import "fmt"

func toMarsWeight(kg float32) float32 {
	return kg * 0.3783
}

func toMarsYears(age float32) float32 {
	return age * 365 / 687
}

func main() {
	myWeightKg := float32(63.0)

	fmt.Printf("My weight on the surface of Mars is %vkg", toMarsWeight(myWeightKg))
	fmt.Printf(" and I would be %v years old.\n", toMarsYears(41))
}
