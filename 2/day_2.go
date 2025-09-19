package main

import (
	"fmt"
)

/*Given an array of integers, return a new array such that each element at
index i of the new array is the product of all the numbers in the original
array except the one at i.

For example, if our input was [1, 2, 3, 4, 5], the expected output would be
[120, 60, 40, 30, 24]. If our input was [3, 2, 1], the expected output would
be [2, 3, 6].

Follow-up: what if you can't use division?*/

func multiplesList(list []int) []int {
	result := make([]int, len(list))
	product := 1

	for _, e := range list {
		product *= e
	}

	for i, e := range list {
		result[i] = product / e
	}

	return result
}

// Brute force. Non optimized.
func multiplesListWithoutDivision(list []int) []int {
	result := make([]int, len(list))

	for i := range list {
		tmp := 1
		for j := range list {
			if i != j {
				tmp *= list[j]
			}
		}
		result[i] = tmp
	}

	return result
}

func main() {
	list := []int{1, 2, 3, 4, 5}
	fmt.Println(multiplesList(list))
	fmt.Println(multiplesListWithoutDivision(list))
}
