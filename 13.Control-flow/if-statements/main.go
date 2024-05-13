package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := 23
	y := 54

	// rand.Intn(a)  => will take the random numbers from 0 to a-1 --> [0,a)
	if z := rand.Intn(1000); z > x {
		fmt.Printf("z = %v is greater than x = %v\n", z, x)
	} else {
		fmt.Printf("z = %v is lesser than x = %v\n", z, x)
	}
	if z := rand.Intn(900); z > y {
		fmt.Printf("z = %v is greater than y = %v\n", z, y)
	} else {
		fmt.Printf("z = %v is lesser than y = %v\n", z, y)
	}
}
