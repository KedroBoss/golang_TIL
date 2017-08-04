package main

import "fmt"

func ex7() {
	sum := 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Println("Sum of all multiples by 3 and 5 are:", sum)
}
