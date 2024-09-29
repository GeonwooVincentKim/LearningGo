package printcomb2

import "fmt"

// PrintComb2 prints all unique combinations of two two-digit numbers in ascending order.
// The output is comma-separated, with the last combination followed by a newline.
// PrintComb2 prints all unique combinations of two two-digit numbers in ascending order.
// The output is comma-separated, with the last combination followed by a newline.
// Output: 00 01, 00 02, 00 03, ..., 00 98, 00 99, 01 02, 01 03, ..., 97 98, 97 99, 98 99
func PrintComb2() {
	for i := 0; i <= 98; i++ {
		for j := i + 1; j <= 99; j++ {
			fmt.Printf("%02d %02d", i, j)
			if !(i == 98 && j == 99) {
				fmt.Print(", ")
			}
		}
	}
	fmt.Print("\n")
}
