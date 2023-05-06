package main

import "fmt"

func vars() (int, int) {
	return 1, 100
}

func main() {
	a, b := vars()
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	_, c := vars()
	fmt.Println("c =", c)
}
