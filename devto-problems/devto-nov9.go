// You are the next contestant on this show, and the host has just showed you the string S. What's the winning last word that you should produce?
// Source - https://code.google.com/codejam/contest/4304486/dashboard#s=p0
package main

import (
	"fmt"
	"strings"
)

func Last(w string) (result []string) {
	l := strings.Split(w, "")
	result = []string{l[0]}

	for _, item := range l[1:] {
		if val := result[0]; item > val {
			result = append([]string{item}, strings.Join(result, ""))
		} else {
			result = append(result, item)
		}
	}
	return
}

func main() {
	fmt.Println(Last("CAB"))
	fmt.Println(Last("JAM"))
	fmt.Println(Last("CODE"))
	fmt.Println(Last("BBAAA"))
	fmt.Println(Last("CABCBBABC"))
	fmt.Println(Last("ABCABCABC"))
	fmt.Println(Last("ZXCASDQWE"))
}
