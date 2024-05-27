package main

import "fmt"

func main() {
	a := map[string]int{
		"sanjay": 21,
		"sugan":  19,
		"Sanju":  23,
	}
	for k, v := range a {
		fmt.Println(k, v)
	}
	// only "key"
	for k := range a {
		fmt.Println(k)
	}

	//only "value"
	for _, v := range a {
		fmt.Println(v)
	}
}
