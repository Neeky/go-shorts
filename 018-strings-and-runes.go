package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const name = "简体中文"
	fmt.Println(name)
	fmt.Println("len(name):", len(name))

	// 如果想以字节方式处理的话，可以按位置迭代
	for i := 0; i < len(name); i++ {
		fmt.Println("name[", i, "]:", name[i])
	}
	fmt.Println()

	chars := utf8.RuneCountInString(name)
	fmt.Println("rune counts ", chars)

	// 如果以字符方式处理的话，可以用 range 迭代
	for index, runeValue := range name {
		fmt.Printf("name[%d]:%c\n", index, runeValue)
	}

}
