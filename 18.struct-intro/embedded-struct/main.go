package main

import "fmt"

func main() {
	type person struct {
		first string
		last  string
		age   int
	}
	type secretAgent struct {
		person
		ltk bool //ltk -- license to kill
	}

	sa1 := secretAgent{
		person: person{
			first: "Sanjay",
			last:  "S",
			age:   21,
		},
		ltk: true,
	}
	p2 := person{
		first: "Sugan",
		last:  "S",
		age:   19,
	}
	fmt.Println(sa1)
	fmt.Println(p2)

	fmt.Printf("%T && %#v\n", sa1, sa1)
	fmt.Printf("%T && %#v\n", p2, p2)

	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)
	fmt.Println(p2.first, p2.last, p2.age)

}
