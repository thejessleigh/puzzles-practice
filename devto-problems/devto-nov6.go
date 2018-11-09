// Calculate the product of all elements in an array
// Source - https://www.codewars.com/kata/product-of-array-items/
package main

import (
	"fmt"
)

func Product(nums []int)(x int) {
	x = 1

	for _, num := range nums {
		x *= num
	}
	return
}


func main(){
	fmt.Println(Product([]int{1, 2, 3, 4, 5}))
	fmt.Println(Product([]int{2, 4, 6, 8, 10}))
}