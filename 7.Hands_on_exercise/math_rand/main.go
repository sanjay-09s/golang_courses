package main

import (
	"fmt"
	"math/rand"
)

func main() {

	/* Intn returns as an int, a non negative pseudo random number in the half
	open interval [0,n) from the default source.
	It panics if n<=0
	*/
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))

}
