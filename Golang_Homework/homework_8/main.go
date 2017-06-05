package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {

	errChannel := make(chan error, 1)
	finished := make(chan bool, 1)
	numCh := make(chan int, 1)

	wg.Add(2)
	for i := 1; i <= 10; i++ {
		go func(errChannel chan<- error, val int) {
			fmt.Println("No", i, " worker")
			if i == 5 {
				errChannel <- errors.New("it's five")
			}
			numCh <- i
		}(errChannel, i)
	}

	go func() {
		wg.Wait()
		finished <- true
	}()

loop:
	for {
		select {
		case <-errChannel:
			fmt.Println("it's 5")
			break loop
		case num := <-numCh:
			fmt.Println(num)
		case <-time.After(100 * time.Millisecond):
			fmt.Println("it's time out")
			break loop

		}
	}
}
