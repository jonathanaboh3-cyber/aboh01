package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	// Handle the special case of the smallest negative number
	if n == -9223372036854775808 {
		z01.PrintRune('-')
		PrintNbr(9)
		PrintNbr(223372036854775808)
		return
	}

	// Handle negative numbers
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	// Recursively print the number
	if n >= 10 {
		PrintNbr(n / 10)
	}
	z01.PrintRune(rune('0' + n%10))
}
