package main

import (
	"ex09/basicatoi2"
	"fmt"
)

func main() {
	fmt.Println(basicatoi2.BasicAtoi2("12345"))
	fmt.Println(basicatoi2.BasicAtoi2("0000000012345"))
	fmt.Println(basicatoi2.BasicAtoi2("012 345"))
	fmt.Println(basicatoi2.BasicAtoi2("Hello World!"))
}
