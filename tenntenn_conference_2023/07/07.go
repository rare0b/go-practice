package main

import "fmt"

func main() {
	var t *Tuple[int, string] = New(10, "hoge")
	fmt.Println(t.X, t.Y)
}

type Tuple[T1, T2 any] struct {
	X T1
	Y T2
}

func New[T1, T2 any](x T1, y T2) *Tuple[T1, T2] {
	return &Tuple[T1, T2]{
		X: x,
		Y: y,
	}
}
