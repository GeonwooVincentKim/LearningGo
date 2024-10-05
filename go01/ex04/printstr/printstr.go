package printstr

import "fmt"

func GetStrLen(s string) int {
	count := 0
	for range s {
		count++
	}
	return count
}

func PrintStr(s string) string {
	// Convert string to slice of runes
	rsStr := []rune(s) // Use []runes to correctly handle multibytes character

	for i := 0; i < GetStrLen(s); i++ {
		fmt.Println("Get S Value - ", rsStr[i])
		fmt.Println("Get S Address - \n", &rsStr[i])
	}
	return s
}

func NewPrintStr(s string) string {
	return s
}
