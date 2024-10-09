package main

import (
	"fmt"
	"statement/newlist"
	// "statement/structure"
)

// func main() {
// 	structure.structure()
// }

// Similar to fields, methods that start with a lowercase letter are unexported (private) in Go
// and cannot be accessed outside their package.
func main() {
	user0 := &newlist.Amazon{}
	user1 := newlist.Amazon{}
	// user0 := newlist.Amazon{UniqueKey: 35}
	// user0.setUniqueKey()

	fmt.Printf("Initializing (User0): %+v\n", user0)
	fmt.Printf("Getting first address (User0): %p\n", &user0)
	user0.SetUniqueKey()
	fmt.Printf("After (User0): %+v\n", user0)
	fmt.Printf("After first address (User0): %p\n", &user0)

	fmt.Println()
	fmt.Printf("Initializing (User1): %+v\n", user1)
	fmt.Printf("Getting first address (User1): %p\n", &user1)
	user1.SetValueUniqueKey()
	fmt.Printf("After (User1): %+v\n", user1)
	fmt.Printf("After first address (User1): %p\n", &user1)
}

// func main() {
// 	v := structure.Vertex{3, 4}
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
