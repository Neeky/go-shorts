package main

import "fmt"

func main() {
	arr := [3]int{1, 10, 100}
	sli := []int{2, 20, 200}

	fmt.Printf("arr's type is %T \n", arr)
	fmt.Printf("sli's type is %T \n", sli)

	sli = append(sli, 2000)
	// arr = append(arr, 1)
	fmt.Printf("sli = %v \n", sli)
}
