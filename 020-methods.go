package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (this Person) SayHello() {
	fmt.Printf("Hello, My name is %s , %d years old.\n", this.name, this.age)
}

// 要设置值的时候要用指针
func (this *Person) SetAge(age int) {
	this.age = age
}

func main() {
	tom := Person{name: "tom"}
	tom.SetAge(16)
	tom.SayHello()
}
