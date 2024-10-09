package main

import "fmt"

// type IntList []int
type IntList []*int

func (list IntList) Print() {
	fmt.Println("List - ", list)
	for i, v := range list {
		if i > 0 {
			fmt.Print(", ")
		}
		if v == nil {
			fmt.Print("nil")
		} else {
			fmt.Print(*v)
		}
	}
	fmt.Println("]")
}

func (list IntList) Append(value int) {
	fmt.Printf("Append before -> %p\n", &list)
	list = append(list, &value)
	fmt.Printf("Append after -> %p\n", &list)
}

// ポインタレシーバ (動的に管理すると)
// func (list *IntList) Append(value int) {
// 	fmt.Printf("Append before -> %p\n", &list)
// 	*list = append(*list, value)
// 	fmt.Printf("Append after -> %p\n", &list)
// }

// func (list IntList) Append(value int) {
// 	fmt.Printf("Append before -> %p\n", &list)
// 	list = append(list, value)
// 	fmt.Printf("Append after -> %p\n", &list)
// }

func (list IntList) Clear() {
	list = nil
}

// 変換されたらならないのに変換されている
func (list *IntList) Double() {
	// for i := range list {
	// 	list[i] *= 2
	// }
	// for _, v := range list {
	// 	if v != nil {
	// 		// Int List(Double関数の方)
	// 		//
	// 		*v *= 2
	// 	}
	// }
	for i := range *list {
		list[i] *= 2
	}
}

func main() {
	a, b, c, d, e := 1, 2, 3, 4, 5
	numbers := IntList{&a, &b, &c, &d, &e}
	// numbers := IntList{1, 2, 3, 4, 5}

	fmt.Println("Initialized - ")
	numbers.Print() // Slice (List)の先頭を指している

}
