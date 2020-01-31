// Structs
// A struct is a collection of fields.

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})

	structA := Vertex{5, 10}

	// Struct fields are accessed using a dot.
	fmt.Println(structA.X, structA.Y)

	// Struct fields can be accessed through a struct pointer.
	// To access the field X of a struct when we have the struct pointer p we could write (*p).X. However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.

	pointA := &structA
	fmt.Println((*pointA).X, pointA.X)  // 5 5 , go语言可以 省略 dereference 直接写作 pointA.X

	println("*****************************")
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
