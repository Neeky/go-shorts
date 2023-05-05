package main

import "fmt"

func main() {
	fmt.Println("hello world")

	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// 如果对应的 key 不存在，map 会返回值类型的默认值，这里是 int 所以返回  0
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	fmt.Println("len(m):", len(m))

	fmt.Println("delete-k2-from m")
	delete(m, "k2")
	fmt.Println("len(m):", len(m))
	fmt.Println(m)

	// 在没有对应的 key 时，标识位会是 false .
	_, prs := m["k2"]
	fmt.Println(prs)

	n := map[string]int{"tom:age": 16, "jerry:age": 18}
	fmt.Println(n)
	fmt.Printf("%T\n", n)

}
