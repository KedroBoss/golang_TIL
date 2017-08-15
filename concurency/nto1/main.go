package main

import (
	"fmt"
	"sync"
)

// n-to-1 is a pattern in which:
// n number of goroutine functions is present
// the function write to a one channel
func main() {
	waitGroup()
	fmt.Println("------------------")
	semaphores()
	// test() //uncomment to see a cool example
}

// channel make -> close
// wg. Add -> Done -> Wait
func waitGroup() {
	chn := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < 10; i++ {
			chn <- i
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			chn <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(chn)
	}()

	for n := range chn {
		fmt.Println(n)
	}
}

func semaphores() {
	n := 50
	c := make(chan int)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func() {
			for d := 0; d < 1000; d++ {
				c <- d
			}
			done <- true
		}()
	}
	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}

// really cool example
// it prints from n to c, but each time the c is decreased
func test() {
	c := make(chan int)
	done := make(chan bool)
	n := 1000

	go func() {
		for i := 0; i < n; i++ {
			c <- i
		}
		close(c)
	}()

	go func() {
		for b := range c {
			for i := n; i >= b; i-- {
				fmt.Println(i)
			}
		}
		done <- true
	}()

	<-done
}
