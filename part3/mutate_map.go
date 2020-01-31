// Insert or update an element in map m:

// m[key] = elem
// Retrieve an element:

// elem = m[key]
// Delete an element:

// delete(m, key)
// Test that a key is present with a two-value assignment:

// elem, ok = m[key]
// If key is in m, ok is true. If not, ok is false.

// If key is not in the map, then elem is the zero value for the map's element type.

// Note: If elem or ok have not yet been declared you could use a short declaration form:

// elem, ok := m[key]

// 使用 ok 判断 map中是否存在某个 key

package main

import "fmt"

//  使用map 取值特别小心， 即使不存在某个 key, m[key] 仍然不报错，仅仅返回 value的数据类型对应的零值，
//	因此应该 使用 v, ok := m[key] 进行 key 是否在 m中的 判定

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}