package main

import (
	"os"

	"github.com/01-edu/z01"
)

type boolean int

const (
	yes boolean = 1
	no  boolean = 0
)

const (
	EvenMsg = "I have an even number of arguments"
	OddMsg  = "I have an odd number of arguments"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func even(nbr int) int {
	if nbr%2 == 0 {
		return 1
	}
	return 0
}

func isEven(nbr int) boolean {
	if even(nbr) == 1 {
		return yes
	}
	return no
}

func main() {
	lengthOfArg := len(os.Args) - 1

	if isEven(lengthOfArg) == yes {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
