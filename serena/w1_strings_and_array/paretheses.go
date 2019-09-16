package main

import (
	"strings"
)

var closers = map[string]string {
	")":"(",
	"}":"{",
	"]":"[",
}

var openers = map[string]string {
	"(":")",
	"{":"}",
	"[":"]",
}



func CheckParentheses(input string) bool {
	var stack []string
	input = strings.ReplaceAll(input, " ", "")

	for _, value := range input {
		value := string(value)
		_, ok := openers[value]

		if ok {
			stack = append(stack, value)
			continue
		}
		val, ok := closers[value]

		if !ok {
			return false
		}
		if len(stack) == 0 {
			return false
		}
		elem := stack[len(stack)-1]
		if elem == val {
			stack = stack[0:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}