package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func DealAPackOfCards(deck []int) {
	// Since we have 12 cards and 4 players, each player gets 3 cards
	for player := 1; player <= 4; player++ {
		fmt.Printf("Player %d: ", player)

		// Calculate the starting and ending index for each player's cards
		start := (player - 1) * 3
		end := start + 3

		// Print the cards for the current player
		for i := start; i < end; i++ {
			if i == end-1 {
				// Last card for the player, no trailing comma
				fmt.Printf("%d", deck[i])
			} else {
				// Cards with trailing comma and space
				fmt.Printf("%d, ", deck[i])
			}
		}

		// Print newline using z01.PrintRune
		z01.PrintRune('\n')
	}
}
