package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	first := true
	for a := 99; a >= 1; a-- {
		for b := a - 1; b >= 0; b-- {
			if !first {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			printNumber(a)
			z01.PrintRune(' ')
			printNumber(b)
			first = false
		}
	}
}

func printNumber(n int) {
	z01.PrintRune(rune(n/10 + '0'))
	z01.PrintRune(rune(n%10 + '0'))
