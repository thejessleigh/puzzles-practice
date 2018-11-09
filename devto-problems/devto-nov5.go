// Write a function to help you calculate whether to guess blue or red when grabbing marbles from a bag
// Arguments
// - blue marbles in the bag at the start
// - red marbles in the bag at the start
// - blue marbles removed so far
// - red marbles removed so far
// Return: probability that you will pull a blue marble next, expressed as a float
// Source - https://www.codewars.com/kata/thinkful-number-drills-blue-and-red-marbles
package main

import (
	"fmt"
)

func Guess(bStart int, rStart int, bGone int, rGone int) (probability float32) {
	var b = bStart - bGone
	var r = rStart - rGone

	probability = float32(b) / float32(r+b)
	return
}

func main() {
	fmt.Println(Guess(5, 5, 2, 3))
	fmt.Println(Guess(50, 100, 30, 85))
}
