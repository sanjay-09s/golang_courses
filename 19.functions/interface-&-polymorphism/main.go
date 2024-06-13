package main

import "fmt"

type person struct {
	first string
}

type secretAgent struct {
	p   person
	ltk bool
}

func (p person) speak() {
	fmt.Println("I am", p.first)
}

func (sa secretAgent) speak() {
	fmt.Println("I am a secret agent", sa.p.first)
}

/*
// polymorphism
func (p person) speak() {
	fmt.Println(p.first)
}

func (sa secretAgent) speak() {
	fmt.Println("I am a secret agent", sa.first)
}
*/

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	sa1 := secretAgent{
		p: person{
			first: "Sanjay",
		},
		ltk: true,
	}
	p2 := person{
		first: "Sugan",
	}

	sa1.speak()
	p2.speak()

	saySomething(sa1)
	saySomething(p2)
}
