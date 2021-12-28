package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) *person {

	p := person{name: name}
	p.age = age
	return &p
}

func main() {

	fmt.Println(person{"Ana", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Carlos", age: 40})

	fmt.Println(newPerson("Paulo", 69))

	s := person{name: "Nina", age: 50}

	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp)
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
