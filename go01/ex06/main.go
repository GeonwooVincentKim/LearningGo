package main

import (
	"ex06/swap"
	"fmt"
)

func main() {
	a := 0
	b := 1

	swap.Swap(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
