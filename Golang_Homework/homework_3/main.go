package main

import (
	"fmt"

	"github.com/go-training/workshops/20170520/homework_3/cart1"
	"github.com/go-training/workshops/20170520/homework_3/cart2"
)

type cart interface {
	Process(...int)
	GetTotal() float32
}

func Output(c cart, price ...int) {
	c.Process(price...)
	fmt.Println(c.GetTotal())
}

func main() {
	buyCart1 := cart1.New("Toy", 1000)
	buyCart2 := cart2.New("Sofa", 2000)

	Output(buyCart1, []int{100, 200, 300}...)
	Output(buyCart2, []int{100, 400, 600}...)
}
