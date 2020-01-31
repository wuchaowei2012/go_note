// Switch evaluation order
// Switch cases evaluate cases from top to bottom, stopping when a case succeeds.

// (For example,

// switch i {
// case 0:
// case f():
// }
// does not call f if i==0.)

// Note: Time in the Go playground always appears to start at 2009-11-10 23:00:00 UTC, a value whose significance is left as an exercise for the reader.


// 从这个例子：
// switch case 可以根据一个变量的值，选择执行 case （constant）后的语句
// 也可以反过来， switch 一个 constant , 把变量放在 case 后

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
