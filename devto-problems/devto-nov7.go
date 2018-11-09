// Implement a difference function, which subtracts one list from another and returns the result.
// Source - https://www.codewars.com/kata/array-dot-diff
package main

import (
	"fmt"
)

func Difference(a []int, b []int) (diff []int) {
	m := make(map[int]bool)

	for _, item := range b {
		m[item] = true
	}

	for _, item := range a {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}
	return
}

func main() {
	var a = []int{1, 2, 3, 4, 5}
	var b = []int{2, 3, 5, 7, 11}
	fmt.Println(Difference(a, b))
	fmt.Println(Difference(b, a))
}
