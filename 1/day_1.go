package main

import (
	"fmt"
)

//Given a list of numbers and a number k, return whether any two numbers from the list add up to k.
//
//For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.
//
//Bonus: Can you do this in one pass?

func addToK(list []int, k int) bool {
	m := make(map[int]bool)

	for _, e := range list {
		if m[e] {
			return true
		}
		m[k-e] = true
	}
	return false
}

func main() {
	var list = []int{10, 15, 3, 7}
	k1 := 16

	fmt.Println(addToK(list, k1))

	k2 := 17
	fmt.Println(addToK(list, k2))
}
