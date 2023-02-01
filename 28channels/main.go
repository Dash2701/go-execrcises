package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("Channels in golang")

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}
	// myCh <- 5
	// fmt.Println(<-myCh)
	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChanelOpen := <-myCh
		fmt.Println(isChanelOpen)
		fmt.Println(val)
		// fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)
	go func(ch chan<- int, wg *sync.WaitGroup) {

		myCh <- 5
		close(myCh)
		// myCh <- 6
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
