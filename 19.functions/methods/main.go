package main

import "fmt"

type person struct {
	first string
}

func (p person) speak() {
	fmt.Println(p.first)
}

func main() {
	p1 := person{
		first: "Sanjay",
	}
	p2 := person{
		first: "Sugan",
	}

	p1.speak()
	p2.speak()
}
