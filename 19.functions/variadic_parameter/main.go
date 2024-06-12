package main

import "fmt"

func main() {
	sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	dump("Sanjay", "Sugan")
	dump("asdf")
	s := []string{"asdf", "ghjk", "jklh", "qwer", "tyui", "oplm"}
	dump("Sanju", s...)
}

//variadic function =>> function can take (any) number of parameters

func sum(i ...int) {
	fmt.Println(i)
	fmt.Printf("%T\n", i)
}

func dump(s string, r ...string) {
	fmt.Println(s)
	fmt.Println(r)
	fmt.Printf("%T\n", r)
}
