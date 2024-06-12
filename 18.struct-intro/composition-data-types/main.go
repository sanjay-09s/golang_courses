package main

import "fmt"

func main() {
	type foo int
	type person struct {
		first string
		last  string
		age   foo
	}

	var f foo = 23
	fmt.Printf("%T\n", f)
	p1 := person{
		first: "Sanjay",
		last:  "S",
		age:   19,
	}

	fmt.Printf("%T\n", p1)
	fmt.Println(f)
	fmt.Printf("%#v\n", p1)
	fmt.Printf("%T\n", p1.age)
}