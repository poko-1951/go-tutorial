package main

import (
	"fmt"
	// "math/rand"
)

func hoge(x, y int) (z, a int) {
	z = x + y
	a = x - y
	return
}

func main() {
	fmt.Println(hoge(2, 3))
}
