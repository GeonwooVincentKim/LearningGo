package newlist

import "fmt"

type IntList []int

func (list IntList) Print() {
	fmt.Print("List Contents : [")
	for i, ele := range list {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(ele)
	}
	fmt.Println("]")
}

func (list *IntList) Append(value int) {
	*list = append(*list, value)
}

func (list *IntList) Truncate() {
	*list = nil
}

// Bad case for Pointer Receiver
func (list IntList) Double() {
	for i, v := range list {
		// if v != nil {
		// 	*v *= 2
		// }
		list[i] = v * 2
	}
}
