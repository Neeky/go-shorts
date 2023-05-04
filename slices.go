package main

import "fmt"

func main() {
	arr := make([]string, 3)
	fmt.Println("arr:", arr)
	fmt.Println("len(arr)", len(arr))

	arr[0] = "a"
	arr[1] = "b"
	arr[2] = "c"
	fmt.Println("set:", arr)
	fmt.Println("get: arr[1]", arr[1])

	arr = append(arr, "d")
	arr = append(arr, "e", "f")
	fmt.Println("append arr :", arr)

	c := make([]string, len(arr))
	copy(c, arr)
	fmt.Println("cpy :", c)

	l := c[:2]
	fmt.Println(l)

	l = c[2:]
	fmt.Print(l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl: ", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerlen := i + 1
		twoD[i] = make([]int, innerlen)

		for j := 0; j < innerlen; j++ {
			twoD[i][j] = j
		}
	}
	fmt.Println(twoD)
}
