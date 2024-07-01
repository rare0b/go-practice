package main

import "fmt"

func Filter[T any](ns []T, f func(i int, v T) bool) []T {
	a := []T{}
	for i, v := range ns {
		if f(i, v) {
			a = append(a, v)
		}
	}
	return a
}

func main() {
	ns := []int{1, 2, 3, 4}
	ms := Filter[int](ns, func(i, v int) bool {
		return v%2 == 0
	})
	fmt.Println(ms) // [2 4]
}
