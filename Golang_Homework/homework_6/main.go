package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	inputList_1 := []string{"1", "3", "5", "7", "9"}
	go func(inputList_1 []string) {
		for _, val := range inputList_1 {
			time.Sleep(300 * time.Millisecond)
			fmt.Println(val)
			//300 600 900 1200 1500
		}
		wg.Done()
	}(inputList_1)

	wg.Add(1)
	inputList_2 := []string{"2", "4", "6", "8", "10"}
	go func(inputList_2 []string) {
		for _, val := range inputList_2 {
			time.Sleep(500 * time.Millisecond)
			fmt.Println(val)
			//500 1000 1500 2000 2500
		}
		wg.Done()
	}(inputList_2)
	wg.Wait()

}
