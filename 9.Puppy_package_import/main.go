package main

import (
	"fmt"

	"github.com/uni-SanjayS/puppy"
)

func main() {
	s1 := puppy.Bark()
	s2 := puppy.Barks()

	fmt.Println(s1)
	fmt.Println(s2)

	// Also we can able to print like this
	fmt.Println(puppy.Bark)
	fmt.Println(puppy.Barks)
}
