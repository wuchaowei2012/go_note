// An interface type is defined as a set of method signatures.

// A value of interface type can hold any value that implements those methods.

// Note: There is an error in the example code on line 22. Vertex (the value type) doesn't implement Abser because the Abs method is defined only on *Vertex (the pointer type).

// 在使用接口时， 某个变量的指针可以调用 值接收者和 指针接收者的方法， 从而可以 将此变量赋值为 接口类型的变量
//               某个变量的值只能调用 值接收者的方法， 从而将此变量赋值为 接口类型的变量

package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	
	// Note: There is an error in the example code on line 22. Vertex (the value type) doesn't implement Abser because the Abs method is defined only on *Vertex (the pointer type).

	// a = v   //cannot use v (type Vertex) as type Abser in assignment: 	Vertex does not implement Abser (Abs method has pointer receiver)

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// ./inference.go:51:6: method redeclared: Vertex.Abs
// func (v Vertex) Abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

