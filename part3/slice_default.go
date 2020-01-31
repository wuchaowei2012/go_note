// Slice defaults
// When slicing, you may omit the high or low bounds to use their defaults instead.

// The default is zero for the low bound and the length of the slice for the high bound.

// For the array

// var a [10]int
// these slice expressions are equivalent:

// a[0:10]
// a[:10]
// a[0:]
// a[:]

package main

import "fmt"

func main() {
	u := []int{2, 3, 5, 7, 11, 13}

	s := u[1:4]
	fmt.Println(s)

	// slice 可以继续切，生成新的slice, 新的slice 底层的 array 仍然是原始的array
	s = s[:2]
	fmt.Println(s)

	s[0] = 0
	fmt.Println(u)

}
