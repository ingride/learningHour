package main

import (
	"fmt"
	"strings"
	"sort"
)


func Duplicate(inputStr string) {
	input := strings.ReplaceAll(inputStr, " ", "")
	count := make(map[string]int)
	var sortedChars []string

	for _, value := range input {
		count[string(value)]++
	}

	for k,v := range count {
		if v > 1 {
			sortedChars = append(sortedChars, k)
		}
	}

	sort.Strings(sortedChars)
	for _, v := range sortedChars {
		fmt.Printf("%s - %d\n",	v, count[v])
	}
}
