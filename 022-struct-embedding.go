package main

import "fmt"

type Address struct {
	city string
}

func (this Address) ShowAddress() {
	fmt.Println(this.city)
}

type Person struct {
	name string
	// 这里一定要与前面定义的 Address 这个 struct 保持一致
	Address
}

func (this Person) Hello() {
	fmt.Printf("hello My name is %s, am from %s \n", this.name, this.Address.city)
}

func main() {
	tom := Person{name: "tom", Address: Address{city: "ShengZheng"}}
	tom.ShowAddress()
	tom.Hello()
}
