package main

import (
	"fmt"
)

func main() {
	var a [3]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1], a[2])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 12}
	fmt.Println(primes)
}
