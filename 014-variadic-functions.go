package main

import "fmt"

func sum(arr ...int) int {
	fmt.Println(arr)
	var total int = 0
	for _, value := range arr {
		total = total + value
	}
	return total
}

func main() {
	total := sum(1, 2, 3, 4)
	fmt.Println("total:", total)

	nums := []int{1, 2, 3}
	total = sum(nums...)
	fmt.Println("total:", total)

}
