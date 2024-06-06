package main

import "fmt"

func main() {
	type person struct {
		first string
		last  string
		age   int
	}

	p1 := person{
		first: "Sanjay",
		last:  "Sengodan",
		age:   21,
	}
	p2 := person{
		first: "Sugan",
		last:  "Sengodan",
		age:   19,
	}
	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Printf("%T && %#v\n", p1, p1)
	fmt.Printf("%T && %#v\n", p2, p2)

	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
}
