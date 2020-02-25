package main

import (
	"container/list"
	"fmt"
)

// list.List of the container/list package is a doubly linked list

func main() {
	// add elements to list
	var intList list.List
	intList.PushBack(11)
	intList.PushBack(12)
	intList.PushBack(13)
	intList.PushBack(14)
	intList.PushBack(15)

	// print elements with neighbours
	for e := intList.Front(); e != nil; e = e.Next() {
		if e.Prev() == nil {
			fmt.Printf("Value: %v Next: %v \n", e.Value, e.Next().Value)
		} else if e.Next() == nil {
			fmt.Printf("Value: %v Prev: %v \n", e.Value, e.Prev().Value)
		} else {
			fmt.Printf("Value: %v Prev: %v Next: %v \n", e.Value, e.Prev().Value, e.Next().Value)
		}
	}
}
