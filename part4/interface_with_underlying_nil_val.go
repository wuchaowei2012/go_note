// Interface values with nil underlying values
// If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.

// In some languages this would trigger a null pointer exception, but in Go it is common to write methods that gracefully handle with a nil receiver (as with the method M in this example.)

// Note that an interface value that holds a nil concrete value is itself non-nil.
// 虽然 value值是 nil, 但是 iterface value自身不是 nil
package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
