package printcombn

import "fmt"

// PrintCombN prints all unique combinations of N different digits in ascending order.
// The output is comma-separated, with the last combination followed by a newline.

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}

	var digits []int
	var printComb func(int, int)

	printComb = func(start, depth int) {
		// Base case: If we've reached the desired depth (n digits)
		if depth == n {
			// Print the current combination
			for i, d := range digits {
				fmt.Printf("%d", d)
				// Add comma and space if it's not the last combination
				if i == n-1 && !(digits[0] == 10-n && digits[n-1] == 9) {
					fmt.Print(", ")
				}
			}
			return
		}

		// Recursive case: Generate combinations
		for i := start; i <= 9; i++ {
			// Add current digit to the combination
			digits = append(digits, i)

			// Recursively generate combinations for the next position
			// The 'start' parameter (i+1) ensures each subsequent digit
			// is greater than the current one, maintaining ascending order
			// and avoiding duplicates
			printComb(i+1, depth+1)

			// Backtrack: remove the last digit to try the next combination
			digits = digits[:len(digits)-1]
		}
	}

	// This recursive approach allows the function to generate all possible
	// combinations of 'n' digits, ensuring they are unique and in ascending order.
	// The 'start' parameter ensures that each digit in a combination is greater
	// than the previous one, avoiding duplicates and maintaining the ascending order.

	printComb(0, 0)
	fmt.Println()
}
