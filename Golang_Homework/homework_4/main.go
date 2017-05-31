package main

import (
	"errors"
	"fmt"
)

func ConverPrint(val interface{}) error {
	if _, ok := val.(string); ok {
		fmt.Println("this is a string")
	} else {
		return errors.New("Error: this is not a string")
	}

	switch val.(type) {
	case string:
		fmt.Println("this is string")
	case int:
		fmt.Println("this is integer")
	}
	return nil
}

func main() {
	fmt.Println(ConverPrint("100"))
}
