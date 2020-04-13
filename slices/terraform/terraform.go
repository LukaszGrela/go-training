package main

import (
	"fmt"
	"strings"
)

func terraform(planets []string) {
	for i := range planets {
		planets[i] = fmt.Sprintf("New %v", planets[i])
	}
}

type Planets []string

func (p Planets) terraform() {
	for i := range p {
		p[i] = fmt.Sprintf("New %v", p[i])
	}
}

func main() {
	planets := Planets{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	fmt.Printf("%s\n", strings.Join(planets, ", "))
	Planets(planets[3:4]).terraform()
	Planets(planets[6:]).terraform()
	fmt.Printf("%s\n", strings.Join(planets, ", "))

}
