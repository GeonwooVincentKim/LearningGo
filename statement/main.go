package main

import (
	"fmt"
	"statement/structure"
)

// func main() {
// 	structure.structure()
// }

func main() {
	v := structure.Vertex{3, 4}
	fmt.Printf("Memory address in Main for v -> %p\n", &v)
	fmt.Println("Memory value in Main for v -> \n", v)
	fmt.Printf("Memory address in Main for v (x)-> %p\n", &v.X)
	fmt.Println("Memory value in Main for v (x) -> \n", v.X)
	fmt.Printf("Memory address in Main for v (y)-> %p\n", &v.Y)
	fmt.Println("Memory value in Main for v (y) -> \n", v.Y)
	fmt.Println("---------------------------------")
	v.Scale(10)

	fmt.Printf("Memory address in After for v -> %p\n", &v)
	fmt.Println("Memory value in After for v -> \n", v)
	fmt.Printf("Memory address in After for v (x)-> %p\n", &v.X)
	fmt.Println("Memory value in After for v (x) -> \n", v.X)
	fmt.Printf("Memory address in After for v (y)-> %p\n", &v.Y)
	fmt.Println("Memory value in After for v (y) -> \n", v.Y)
	a := 150
	fmt.Printf("Type is -> %T\n", a)
	fmt.Println(v.Abs())
}
