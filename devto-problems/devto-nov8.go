// Construct a function that accepts two string and returns true if a portion of st1 characters can be arranged to match str2, otherwise false
// Source - https://www.codewars.com/kata/scramblies
package main

import (
	"fmt"
	"strings"
)

func Scramble(str1 string, str2 string) (result bool) {
	m := make(map[string]int)

	for _, item := range strings.Split(str1, "") {
		m[item] += 1
	}

	for _, item := range strings.Split(str2, "") {
		if val, ok := m[item]; !ok || val == 0 {
			return false
		} else {
			m[item] -= 1
		}
	}

	return true
}

func main() {
	fmt.Println(Scramble("rkqodlw", "world"))
	fmt.Println(Scramble("cedewaraaaossoqqyt", "codewars"))
	fmt.Println(Scramble("katas", "steak"))
}
