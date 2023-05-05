package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, item := range nums {
		sum = sum + item
	}

	fmt.Println("sum:", sum)

	for index, num := range nums {
		if num == 3 {
			fmt.Println("index: ", index)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("k:%s v:%s \n", k, v)
	}

	for k := range kvs {
		fmt.Printf("k:%s \n", k)
	}

	for i, c := range "AB" {
		fmt.Printf("i: %d, c: %c \n", i, c)
	}
}
