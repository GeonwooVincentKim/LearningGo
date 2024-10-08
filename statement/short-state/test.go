package main

import (
	"fmt"
	"math"
	"reflect"
)

func pow(x, n, lim float64) float64 {
	reflect.TypeOf(x)
	reflect.TypeOf(n)
	fmt.Println(reflect.TypeOf(lim))
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
