package main

import (
	"fmt"
	"sync"
)

var waitgroup sync.WaitGroup

func PrintStrint(num int) {
	fmt.Println(num)
	waitgroup.Done()
}

func main() {
	for i := 0; i <= 5; i++ {
		waitgroup.Add(1)
		go PrintStrint(i)
	}

	waitgroup.Wait()
	fmt.Println("done!")
}
