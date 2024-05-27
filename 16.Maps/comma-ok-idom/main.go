package main

import "fmt"

func main() {
	a := map[string]int{
		"Sanjay": 21,
		"Sugan":  19,
		"Sanju":  23,
	}
	// "ok" -- idom
	// if the key exist for the key then "true" else "false"
	fmt.Println("-----------------")
	fmt.Println(a)
	v, ok := a["mani"]
	if ok {
		fmt.Println("Key exixt and value is ", v)
	} else {
		fmt.Println("Key doesn't exist")
	}

	fmt.Println("------------------")
	fmt.Println("Another way to check, new syntax")
	if k, ok := a["mani"]; ok {
		fmt.Println("Key exixt and value is ", k)
	} else {
		fmt.Println("Key doesn't exist")
	}

	fmt.Println("-------------------")
	a["mani"] = 44
	fmt.Println("After adding the new key \"mani\"")
	if k, ok := a["mani"]; ok {
		fmt.Println("Key exixt and value is ", k)
	} else {
		fmt.Println("Key doesn't exist")
	}
	fmt.Println("-------------------")
}
