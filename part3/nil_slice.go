// Nil slices
// The zero value of a slice is nil.

// A nil slice has a length and capacity of 0 and has no underlying array.

// nil slice的定义， length capacity = 0 && 没有底层的 array

package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}
