// package main

// func swap(x, y string) (string, string) {
// 	return y, x
// }

// func swap2(x, y int) (int, int) {
// 	return y, x
// }

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
	// a, b := swap("hello", "world")
	// fmt.Println(a, b)

	// a, b := swap("")
	// name := "test" // short-hand way, go gives you a type
	// var name string = "test"
}
