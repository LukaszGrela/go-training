package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("Listing 9.7")
	message := "uv vagreangvbany fcnpr fgngvba"

	for i := 0; i < len(message); i++ {
		c := message[i]
		if c >= 'a' && c <= 'z' {
			c = c + 13
			if c > 'z' {
				c = c - 26
			}
		}
		fmt.Printf("%c", c)
	}

	fmt.Println("\nListing 9.8")
	question := "Co słychać?"
	fmt.Println(len(question), "bytes")
	fmt.Println(utf8.RuneCountInString(question), "runes")
	c, size := utf8.DecodeRuneInString(question)
	fmt.Printf("First rune: %c %v bytes", c, size)

	fmt.Println("\nListing 9.9")
	for _, c := range question {
		fmt.Printf("%c ", c)
	}
	fmt.Println("")

	decodeCeasar()

	fmt.Println("CSOITEUIWUIZNSROCNKFD")
	fmt.Println(decodeVigenere("CSOITEUIWUIZNSROCNKFD", "GOLANG"))
	fmt.Println(decodeVigenere(encodeVigenere(decodeVigenere("CSOITEUIWUIZNSROCNKFD", "GOLANG"), "GOLANG"), "GOLANG"))
	fmt.Println(encodeVigenere("LUKASZ", "AXB"))
	fmt.Println(decodeVigenere(encodeVigenere("LUKASZ", "AXB"), "AXB"))
}

func decodeCeasar() {
	quote := "L FDPH, L vdz, L frqtxhuhg."

	fmt.Println("Julius Caesar Quote")
	for i := 0; i < len(quote); i++ {
		c := quote[i]
		if c >= 'a' && c <= 'z' {
			c = c - 3
			if c > 'z' {
				c = c - 26
			}
		}
		if c >= 'A' && c <= 'Z' {
			c = c - 3
			if c > 'Z' {
				c = c - 26
			}
		}
		fmt.Printf("%c", c)
	}
}

func decodeVigenere(encoded string, keyword string) string {
	decoded := ""
	for i := 0; i < len(encoded); i++ {
		c := encoded[i] - 'A'
		k := keyword[i%len(keyword)] - 'A'
		decoded += string(((c - k + 26) % 26) + 'A')
	}
	return decoded
}

func encodeVigenere(input string, keyword string) string {
	encoded := ""
	for i := 0; i < len(input); i++ {
		c := input[i] - 'A'
		k := keyword[i%len(keyword)] - 'A'
		encoded += string(((c + k + 26) % 26) + 'A')
	}
	return encoded

}
