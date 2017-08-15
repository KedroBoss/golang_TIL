package main

import (
	"fmt"
)

func main() {
	c := factorial(8)
	for n := range c {
		fmt.Println(n)
	}
}

func factorial(number int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for i := number; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}
