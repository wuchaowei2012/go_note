// The equivalent thing happens in the reverse direction.

// Functions that take a value argument must take a value of that specific type:

// var v Vertex
// fmt.Println(AbsFunc(v))  // OK
// fmt.Println(AbsFunc(&v)) // Compile error!
// while methods with value receivers take either a value or a pointer as the receiver when they are called:

// 一般函数的参数不支持 变量转化

// var v Vertex
// fmt.Println(v.Abs()) // OK
// p := &v
// fmt.Println(p.Abs()) // OK
// In this case, the method call p.Abs() is interpreted as (*p).Abs().

// method 支持变量转化

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}
// 使用 value 作为接收者，仅仅拷贝原始变量的一个副本
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}
