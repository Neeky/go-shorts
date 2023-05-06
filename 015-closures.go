package main

import "fmt"

func Seqint() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := Seqint()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	ni := Seqint()
	fmt.Println(ni())
	fmt.Println(nextInt())
}
