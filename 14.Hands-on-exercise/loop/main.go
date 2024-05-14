package main

import "fmt"

func main() {

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("Outer loop %v \t Innerloop %v\n", i, j)
		}
		fmt.Println("")
	}

}
