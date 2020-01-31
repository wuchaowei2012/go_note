// Switch
// A switch statement is a shorter way to write a sequence of if - else statements. It runs the first case whose value is equal to the condition expression.

// Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs the selected case, not all the cases that follow. In effect, the break statement that is needed at the end of each case in those languages is provided automatically in Go. Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.

// switch case 不限制为 constant 也不限制为 integers
// switch case 根据 var 的值，选择执行不同的case。 也可以使用类似 if语句中的 临时变量

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {  //os 为类似于 if 中的临时变量
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	// fmt.Println(os) //undeclared name: oscompiler

}


func main1() {
	os := runtime.GOOS
	fmt.Println("runtime.goos:", os)

	fmt.Print("Go runs on ")
	switch os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
