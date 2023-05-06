package main

import "fmt"

func fact(n int) int {
	if n <= 1 {
		return 1
	} else {
		return n * fact(n-1)
	}
}

func main() {
	var n int = 3
	fmt.Println(fact(n))

	var fib func(int) int
	fib = func(n int) int {
		if n <= 2 {
			return 1
		} else {
			return fib(n-1) + fib(n-2)
		}
	}

	fmt.Println(fib(5))
}
