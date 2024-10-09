package newlist

import "fmt"

type Amazon struct {
	UniqueKey int
}

// Pointer Receiver
func (a *Amazon) SetUniqueKey() {
	fmt.Printf("Getting Function Address (Before)-> %p\n", &a)
	fmt.Println("Getting Function Value (Before) -> ", a.UniqueKey)
	a.UniqueKey++
	fmt.Printf("Getting Function Address (After)-> %p\n", &a)
	fmt.Println("Getting Function Value (After) -> ", a.UniqueKey)
}

// Value Receiver
func (a Amazon) SetValueUniqueKey() {
	fmt.Printf("Getting Function Address (Before)-> %p\n", &a)
	fmt.Println("Getting Function Value (Before) -> ", a.UniqueKey)
	a.UniqueKey++
	fmt.Printf("Getting Function Address (After)-> %p\n", &a)
	fmt.Println("Getting Function Value (After) -> ", a.UniqueKey)
}

// func (a *Amazon) getString() string {
// 	return "Amazon"
// }
