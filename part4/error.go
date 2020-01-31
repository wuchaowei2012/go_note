// Go programs express error state with error values.

// The error type is a built-in interface similar to fmt.Stringer:

// type error interface {
//     Error() string
// }
// (As with fmt.Stringer, the fmt package looks for the error interface when printing values.) 
// 当打印 error 接口值时， fmt 包寻找 error 接口的 Error 方法

// Functions often return an error value, and calling code should handle errors by testing whether the error equals nil.

// i, err := strconv.Atoi("42")
// if err != nil {
//     fmt.Printf("couldn't convert number: %v\n", err)
//     return
// }
// fmt.Println("Converted integer:", i)
// A nil error denotes success; a non-nil error denotes failure.


package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)   //%v  value     %s string
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)   // // 当打印 error 接口值时， fmt 包寻找 error 接口的 Error 方法
	}
}
