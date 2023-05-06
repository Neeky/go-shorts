package main

import "fmt"

func zeroval(i int) {
	i = 0
}

func zeroptr(i *int) {
	*i = 0
}

func main() {
	i := 1
	fmt.Println("i:", i)

	zeroval(i)
	fmt.Println("i:", i)

	zeroptr(&i)
	fmt.Println("i:", i)
}
