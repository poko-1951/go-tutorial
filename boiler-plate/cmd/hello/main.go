package main

import (
	"fmt"
	// "math/rand"
)

const foo = "foo!!"

func hoge(x, y int) (z, a int) {
	z = x + y
	a = x - y
	return
}

var c, python bool = true, true

func main() {
	// i := 2
	// j := int64(i)
	// fmt.Println(hoge(2, 3))
	// fmt.Println(i, j, c, python)
	fmt.Println(foo)
}
