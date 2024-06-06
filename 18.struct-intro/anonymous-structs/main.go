package main

import "fmt"

func main() {
	type person struct {
		first string
		last  string
		age   int
	}

	// anonymous struct
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "Sanjay",
		last:  "S",
		age:   21,
	}

	p2 := person{
		first: "Sugan",
		last:  "S",
		age:   19,
	}

	fmt.Printf("%T\n", p1)
	fmt.Printf("%T\n", p2)
}
