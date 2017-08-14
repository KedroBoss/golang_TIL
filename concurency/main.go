package main

import "fmt"
import "sync"
import "time"
import "runtime" // for parallelism

var wg sync.WaitGroup
var counter int
var mutex sync.Mutex

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Without ths it'll be concurrent not parallel
}

func main() {
	// wg.Add(2)
	// go funcOne()
	// go funcTwo()

	// wg.Wait()
	channel()
}

func funcOne() {
	for i := 0; i < 10; i++ {
		fmt.Println("One:", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
		mutex.Lock()
		x := counter
		x++
		counter = x
		mutex.Unlock()
		fmt.Println(counter)
	}
	wg.Done()
}

func funcTwo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Two:", i)
		time.Sleep(time.Duration(10 * time.Millisecond))
		mutex.Lock()
		x := counter
		x++
		counter = x
		mutex.Unlock()
		fmt.Println(counter)
	}
	wg.Done()
}

func channel() {
	chn := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			chn <- i
		}
	}()

	go func() {
		for {
			fmt.Println(<-chn)
		}
	}()
	time.Sleep(time.Second)
}
