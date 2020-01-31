// Creating a slice with make
// Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.

// The make function allocates a zeroed array and returns a slice that refers to that array:
// 第2个元素定义 len, 第3个元素定义 cap
// a := make([]int, 5)  // len(a)=5

// To specify a capacity, pass a third argument to make:

// b := make([]int, 0, 5) // len(b)=0, cap(b)=5

// len 是 slice 中包含的元素个数， cap 是 slice 中第一个元素到 底层array 最后一个元素间的元素个数

// b = b[:cap(b)] // len(b)=5, cap(b)=5
// b = b[1:]      // len(b)=4, cap(b)=4

package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
