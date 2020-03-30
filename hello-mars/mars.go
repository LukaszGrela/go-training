package main

import "fmt"

func toMarsWeight(kg float32) float32 {
	return kg * 0.3783
}

func main() {
	myWeightKg := float32(63.0)

	fmt.Print("My weigt on the surface of Mars is ")
	fmt.Print(toMarsWeight(myWeightKg))
	fmt.Println("kg.")
}
