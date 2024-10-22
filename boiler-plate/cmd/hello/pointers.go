package main

import (
	"fmt"
)

func main() {
	i, j := 42, 2701

	l := i
	fmt.Println(l)
	l = 1
	fmt.Println(l)
	fmt.Println(i)

	l = j
	fmt.Println(l)

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}
