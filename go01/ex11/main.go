package main

import (
	"ex11/sortintegertable"
	"fmt"
)

func main() {
	s := []int{5, 4, 3, 2, 1, 0}
	sortintegertable.SortIntegerTable(s)
	fmt.Println(s)
}
