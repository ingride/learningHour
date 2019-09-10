package main

import (
	"fmt"
	"unicode"
)

func calculateCharacters(s string, m map[rune]int ) {
	for _, c := range s {
		if !unicode.IsLetter(c) {
			continue
		}
		if val, ok := m[c]; ok {
			m[c] = val + 1
		} else {
			m[c] = 1
		}
	}
}

func printDups(m map[rune]int) {
	for k,v := range m {
		if (v>1) {
			fmt.Println(string(k), " => ", v)
		}
	}
}

func main() {

	var m map[rune]int

	m = make(map[rune]int)

	calculateCharacters("ingrid and serena learn go", m)

	printDups(m)

}