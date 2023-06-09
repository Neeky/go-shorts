package main

import "fmt"

func main() {
	var arr [5]int
	fmt.Println(arr)

	arr[4] = 100
	fmt.Println("set:", arr)
	fmt.Println("get:", arr[4])
	fmt.Println("len(arr):", len(arr))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD:", twoD)
}
