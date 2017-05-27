package main

import (
	"fmt"

	"github.com/go-training/workshops/20170520/homework_2/car"
)

func main() {
	car1 := car.New("Honda", 100, "Grey")

	fmt.Println("Car 1's name:", car1.GetName())
	fmt.Println(car1.Add(250, 120))
	fmt.Println(car1.Remove(80))
	fmt.Println(car1.UpdateColor("Yellow"))
}
