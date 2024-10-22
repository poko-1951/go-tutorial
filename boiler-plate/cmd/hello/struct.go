package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

type Hoge struct {
	X, Y int
}

var (
	v1 = Hoge{1, 2}
	v2 = Hoge{X: 1}
	v3 = Hoge{}
	p = &Hoge{2, 2}
)

func main() {
	// v := Vertex{1, 2}
	// v.X = 4
	// p := &v
	// p.X = 1e9
	// fmt.Println(v)
	fmt.Println(v2, p, v2, v3)
}
