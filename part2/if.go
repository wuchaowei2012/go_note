// If
// Go's if statements are like its for loops; the expression need not be surrounded by parentheses ( ) but the braces { } are required.

// if 判定语句不需要 括号（）， 但是 运行语句需要 {} 括起来


package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))  // 将 float 转化为 string
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
