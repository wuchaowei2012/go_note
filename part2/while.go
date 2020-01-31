// At that point you can drop the semicolons: C's while is spelled for in Go.
// 在 go 语言中，while 被拼写为 for

package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

