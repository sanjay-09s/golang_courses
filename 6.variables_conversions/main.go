package main

import "fmt"

func main() {
	y := 42
	x := 42.0
	fmt.Printf("%v of type %T\n", y, y)
	fmt.Printf("%v of type %T\n", x, x)

	var m float32 = 43.724
	fmt.Printf("%v of type %T\n", m, m)

	/*
		 	THis does not work

			In go you can't take a value that is float32
			and store the data in variable that is declares as float64
			z = m
			fmt.Printf("%v of type %T",z,z)

	*/

	// This doe's works

	x = float64(m)
	fmt.Printf("%v of type %T\n", x, x)
}
