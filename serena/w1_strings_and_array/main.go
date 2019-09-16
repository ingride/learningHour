package main

import (
	"os"
)
func main() {
	runCommand(os.Args)
}

func runCommand(args []string) {
	command := args[1]
	switch command {
	case "parens":
		result := CheckParentheses(args[2])
		if result {
			println("This is a balanced string")
		} else {
			println("This is an imbalanced string")
		}
	case "duplicate":
		Duplicate(args[2])
	}

}