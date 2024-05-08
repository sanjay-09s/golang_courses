package main

import "fmt"

func main() {
	a := 0
	b, c, _, d := 1, 2, 3, "Sanjay"
	fmt.Println(a, b, c, d)

	//This doesn't work
	/*
		b,c,e := 5,6,7
		fmt.Println(b,c)    ---this gives error [ e is declared not used ]
	*/
	var g int
	fmt.Println(g)
	g = 24
	fmt.Println(g)
	g = 35
	fmt.Println(g)
}
