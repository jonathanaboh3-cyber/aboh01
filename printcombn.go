package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}

	// Create a slice to store current combination
	comb := make([]rune, n)
	generateComb(comb, 0, '0', n)
	z01.PrintRune('\n')
}

func generateComb(comb []rune, pos int, start rune, n int) {
	if pos == n {
		// Print the combination
		for i := 0; i < n; i++ {
			z01.PrintRune(comb[i])
		}

		// Check if this is the last combination
		// Last combination is when we have the highest possible digits
		// For n=3, last is "789", for n=2, last is "89", etc.
		isLast := true
		for i := 0; i < n; i++ {
			if comb[i] != '9'-rune(n-i-1) {
				isLast = false
				break
			}
		}

		if !isLast {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		return
	}

	// Generate all possible digits for current position
	// The maximum digit we can use at position i is '9' - (n - i - 1)
	for digit := start; digit <= '9'-rune(n-pos-1); digit++ {
		comb[pos] = digit
		generateComb(comb, pos+1, digit+1, n)
	}
}
