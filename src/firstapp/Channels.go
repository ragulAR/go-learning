package main

import (
	"fmt"
	"sync"
)

var waitGroup = sync.WaitGroup{}

func main() {

	ch := make(chan int, 50) // introduce internal data store
	waitGroup.Add(2)

	go func(ch <-chan int) { // receives messages from channel
		//for {
		//	if i, ok := <- ch; ok { // checks if channel is closed
		//		fmt.Println(i)
		//	} else {
		//		break
		//	}
		//}
		for i := range ch {
			fmt.Println(i)
		}
		waitGroup.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 99 // sends message into the channel
		ch <- 69
		close(ch) // close channel
		// ch <- 88 // will panic as the channel is closed
		waitGroup.Done()
	}(ch)
	waitGroup.Wait()
}
