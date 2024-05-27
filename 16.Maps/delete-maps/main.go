package main

import "fmt"

func main() {
	a := map[string]int{
		"Sanjay": 21,
		"Sugan":  19,
		"Sanju":  23,
	}
	fmt.Println("a")
	fmt.Printf("%#v\n", a)
	fmt.Println("---------------------")
	delete(a, "Sugan")
	fmt.Println("After deleting")
	fmt.Println("a")
	fmt.Printf("%#v\n", a)
	fmt.Println("---------------------")
	fmt.Println("Deleting an element that doesn't exist")
	delete(a, "SanjayS")
	fmt.Println("Printing an element that doesn't exist")
	fmt.Println(a["main"])
	fmt.Println("----------------------")
	fmt.Println("Adding a new element")
	a["Mani"] = 42
	fmt.Println(a)
	fmt.Printf("%#v\n", a)
}
