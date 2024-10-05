package main

import (
	"ex10/atoi"
	"fmt"
)

func main() {
	fmt.Println(atoi.Atoi("12345"))
	fmt.Println(atoi.Atoi("0000000012345"))
	fmt.Println(atoi.Atoi("012 345"))
	fmt.Println(atoi.Atoi("Hello World!"))
	fmt.Println(atoi.Atoi("+1234"))
	fmt.Println(atoi.Atoi("-1234"))
	fmt.Println(atoi.Atoi("++1234"))
	fmt.Println(atoi.Atoi("--1234"))
}
