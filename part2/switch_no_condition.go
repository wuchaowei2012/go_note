// Switch with no condition
// Switch without a condition is the same as switch true.

// This construct can be a clean way to write long if-then-elseif-else chains.
// 此时 switch case 执行第一个满足的 case 语句
package main

import (
	"fmt"
	// "time"
)

// func main() {
// 	t := time.Now()
// 	switch {
// 	case t.Hour() < 12:
// 		fmt.Println("Good morning!")
// 	case t.Hour() < 17:
// 		fmt.Println("Good afternoon.")
// 	default:
// 		fmt.Println("Good evening.")
// 	}
// }


func main() {
	t := 20
	switch {
	case t < 12:
		fmt.Println("Good morning!")
	case t < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

