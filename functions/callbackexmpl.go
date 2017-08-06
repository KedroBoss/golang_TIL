package main

import "fmt"

// takes: slice of integers, function takes integers and returns booleans
// returns slice of integers
func filter(numbers []int, callback func(int) bool) []int {
	var slice []int
	for _, n := range numbers {
		if callback(n) {
			slice = append(slice, n)
		}
	}
	return slice
}

func filterOut() {
	slice := filter([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, func(n int) bool {
		return n > 1
	})
	fmt.Println(slice)
}
