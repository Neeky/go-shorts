package main

import "fmt"

type Person struct {
	name string
	age  int
}

func newPerson(name string) *Person {
	person := Person{name: name}
	person.age = 100
	return &person
}

func main() {
	fmt.Println(Person{name: "bob", age: 16})
	fmt.Println(Person{name: "jerry"})

	fmt.Println(&Person{name: "leo"})
	fmt.Println(newPerson("leo"))

	s := &Person{name: "banana", age: 16}
	fmt.Println(s.age)
	s.age = 100
	fmt.Println(s.age)

}
