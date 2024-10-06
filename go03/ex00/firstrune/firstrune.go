package firstrune

import "fmt"

func FirstRune(s string) rune {
	for _, value := range s {
		fmt.Println("Get Value -> ", value)
		return value
	}
	return 0
}
