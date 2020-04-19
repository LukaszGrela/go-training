package main

import (
	"fmt"
)

func main() {
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dwarfs = append(dwarfs, "Orcus")
	fmt.Println(dwarfs)

	newSlice := make([]int, 0, 4)
	for step := 0; step < 10000; step++ {
		capacity := cap(newSlice)
		newSlice = append(newSlice, step)
		if capacity != cap(newSlice) {
			fmt.Printf("%v\n", cap(newSlice))
		}
	}
}
