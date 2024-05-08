package main

import "fmt"

func main() {
	const name, age = "Sanjay", 21
	fmt.Printf("%v is %v years old.\tType of name is \"%T\" and age is \"%T\"", name, age, name, age)
}
