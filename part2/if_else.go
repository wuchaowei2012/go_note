// If and else
// short statement declared inside an if are also available inside any of the else blocks.

// (Both calls to pow return their results before the call to fmt.Println in main begins.)
// main 函数中，先执行 两个pow， 然后再执行 main函数里的 fmt.Println 函数

package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		// 仅仅输出 提示， 对程序运行没有什么作用
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}