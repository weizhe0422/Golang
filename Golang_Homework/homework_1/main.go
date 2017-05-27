package main

import (
	"fmt"
)

func calc(start int, end int) (int, float32) {
	sum := 0
	count := 0

	for index := start; index <= end; index++ {
		//if index%2 == 0 {
		//	continue
		//}
		sum += index
		count++
	}

	return sum, float32(sum) / float32(count)

}

func calc2(nums ...int) (int, float32) {
	sum := 0
	count := 0

	for _, val := range nums {
		if val%2 == 0 {
			continue
		}
		sum += val
		count++
	}
	return sum, float32(sum) / float32(count)
}

func main() {
	fmt.Println(calc(1, 100))
	fmt.Println(calc(50, 100))
	fmt.Println(calc2(1, 2, 3, 4, 5, 6))
	fmt.Println(calc2(101, 202, 303))
}
