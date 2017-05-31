package main

import (
	"fmt"
)

type MyValError struct {
	val          int
	errorMessage string
}

func (e MyValError) Error() string {
	return fmt.Sprintf("Error: val: %d, message: %s", e.val, e.errorMessage)
}


func PrintVal(inputValue int) (message, error) {
	if inputValue < 10 {
		return "", MyValError{
			val:          inputValue,
			errorMessage: "this value less than 10",
		}
	}

	return "this value large than 10", nil
}

func main() {
	fmt.Println(PrintVal(19))
}
