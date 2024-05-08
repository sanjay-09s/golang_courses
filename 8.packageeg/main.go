package main

import (
	"fmt"
	"golang_course/8.packageeg/furtherexplored"
)

var x = 42

func main() {
	x := 34
	fmt.Println(x)

	PrintMe()

	y := 43
	fmt.Println(y)

	p := person{
		"Sanjay",
	}
	p.sayHello()

	if z := 50; z < 100 {
		fmt.Println("Z is here")
	}

	// z is not accessible here

	//packages in the scope eg :: furtherexplored
	furtherexplored.Fascinating()
	fmt.Println(furtherexplored.Planetspeed)
}

func PrintMe() {
	fmt.Println(x)
}

type person struct {
	name string
}

func (p person) sayHello() {
	fmt.Println("Hi my name is ", p.name)
}
