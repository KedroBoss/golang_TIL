package main

import (
	"fmt"
)

func main() {
	ex2 := ex1
	fmt.Println(ex2(4))
	fmt.Println(ex3(1, 4, 2, 7, 9))
}

func ex1(number int) (int, bool) {
	return number / 2, number%2 == 0
}

func ex3(numbers ...int) int {
	var x int
	for _, v := range numbers {
		if v > x {
			x = v
		}
	}
	return x
}
