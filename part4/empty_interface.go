// The empty interface
// The interface type that specifies zero methods is known as the empty interface:

// 空接口是没有指定任何方法的接口

// interface{}
// An empty interface may hold values of any type. (Every type implements at least zero methods.)

// Empty interfaces are used by code that handles values of unknown type. For example, fmt.Print takes any number of arguments of type interface{}.

package main

import "fmt"

func main() {
	// 定义一个 nil interface
	var i interface{}
	describe(i)  //(<nil>, <nil>)

	i = 42
	describe(i)   //(42, int)

	i = "hello"
	describe(i) //(hello, string) 打印出 接口真实的数据类型
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
