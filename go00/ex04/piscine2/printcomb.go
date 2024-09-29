package piscine2

import "fmt"

// PrintComb result should be like this: 012, 013, 014, 015, 016, 017, 018, 019, 023, ...., 689, 789, 89
// PrintComb prints all unique combinations of three different digits in ascending order.
// The output is comma-separated, with the last combination followed by a newline.
func PrintComb() {
	for i := 0; i <= 7; i++ {
		for j := i + 1; j <= 8; j++ {
			for k := j + 1; k <= 9; k++ {
				fmt.Printf("%d%d%d", i, j, k)
				if i != 7 || j != 8 || k != 9 {
					fmt.Print(", ")
				}
			}
		}
	}
	fmt.Println()
}
