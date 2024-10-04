package main

import (
	"ex00/pointone"
	"fmt"
)

func main() {
	// i, j := 42, 2701

	// fmt.Println(i)
	// fmt.Println(j)
	// fmt.Println(&i)
	// fmt.Println(&j)

	// var p *int
	// fmt.Println(&p)

	// p = &i
	// fmt.Println(&i)
	// fmt.Println(p)

	n := 0
	fmt.Println(n)
	fmt.Println(&n)
	fmt.Println("")
	pointone.PointOne(&n)
	fmt.Println(n)
}
