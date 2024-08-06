package main

import (
	"errors"
	"fmt"
)

func main() {
	var a, b int = 5, 3
	result, remainder, err := divideNumbers(a, b)

	switch {
	case err != nil:
		fmt.Printf(err.Error())
	case remainder == 0:
		fmt.Printf("the result is %v", result)
	default:

		fmt.Printf("the result is %v and remaminder is %v", result, remainder)
	}
}

func divideNumbers(a int, b int) (int, int, error) {
	result := a / b
	remainder := a % b

	var err error

	if b == 0 {
		err = errors.New("sopu")
		return 0, 0, err
	}

	return result, remainder, err
}
