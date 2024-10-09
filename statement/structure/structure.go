package structure

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
	fmt.Printf("Memory address of v in Scale: %p\n", &v)
	fmt.Println("Memory value of v in Scale: \n", v)
	fmt.Printf("Memory address in v (x)-> %p\n", &v.X)
	fmt.Println("Memory value in v (x) -> \n", v.X)
	fmt.Printf("Memory address in v (y)-> %p\n", &v.Y)
	fmt.Println("Memory value in v (y) -> \n", v.Y)
	fmt.Println("---------------------------------")

	v.X = v.X * f
	v.Y = v.Y * f
}

// func main() {
// 	v := Vertex{3, 4}
// 	fmt.Printf("Memory address in Main for v -> %p\n", &v)
// 	fmt.Println("Memory value in Main for v -> \n", v)
// 	fmt.Printf("Memory address in Main for v (x)-> %p\n", &v.X)
// 	fmt.Println("Memory value in Main for v (x) -> \n", v.X)
// 	fmt.Printf("Memory address in Main for v (y)-> %p\n", &v.Y)
// 	fmt.Println("Memory value in Main for v (y) -> \n", v.Y)
// 	fmt.Println("---------------------------------")
// 	v.Scale(10)

// 	fmt.Printf("Memory address in After for v -> %p\n", &v)
// 	fmt.Println("Memory value in After for v -> \n", v)
// 	fmt.Printf("Memory address in After for v (x)-> %p\n", &v.X)
// 	fmt.Println("Memory value in After for v (x) -> \n", v.X)
// 	fmt.Printf("Memory address in After for v (y)-> %p\n", &v.Y)
// 	fmt.Println("Memory value in After for v (y) -> \n", v.Y)
// 	a := 150
// 	fmt.Printf("Type is -> %T\n", a)
// 	fmt.Println(v.Abs())
// }
