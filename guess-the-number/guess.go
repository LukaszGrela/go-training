package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Guess the number I'm thinking (0 .. 100)")

	theNumber := rand.Intn(100)

	guessed := math.MaxInt8

	for guessed != theNumber {
		fmt.Printf("What is your guess=")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if isValidEntry(text) {
			// a number
			c, _ := strconv.Atoi(text)
			guessed = c
			if guessed > theNumber {
				fmt.Println("My number is smaller than yours...")
			} else if guessed < theNumber {
				fmt.Println("My number is bigger than yours...")
			}
		} else {
			fmt.Println("Didn't get that...")
		}
	}
	fmt.Println("You've guessed it...")
}

func isValidEntry(input string) bool {
	re, _ := regexp.Compile(`[^\d]+`)
	return len(re.FindAllString(input, -1)) == 0
}
