package main

import (
	"errors"
	"fmt"
)

func fa(i int) (int, error) {
	if i == 0 {
		return -1, errors.New("0 is not supported .\n")
	}
	return i + 1, nil
}

func main() {
	res, err := fa(10)
	if err != nil {
		fmt.Printf("failed error = %s \n", err)
	} else {
		print("result = ", res, "\n")
	}

	res, err = fa(0)
	if err != nil {
		fmt.Printf("failed error = %s \n", err)
	} else {
		print("result = ", res, "\n")
	}
}
