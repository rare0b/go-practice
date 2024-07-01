package main

import "fmt"

func Apply[T any](a []T, f func(i int, v T)) {
	for i, v := range a {
		f(i, v)
	}
}

func main() {
	var sum int
	Apply[int]([]int{10, 20}, func(i, v int) {
		sum += v
	})
	fmt.Println(sum) // 30
}
