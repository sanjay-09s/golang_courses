package main

import (
	"fmt"
	"strconv"
)

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprint("I am reading the book with title", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("I am at the page", strconv.Itoa(int(c)))
}

func main() {
	b := book{
		title: "Steal Like An Artist",
	}
	var c count = 32
	fmt.Println(b)
	fmt.Println(c)
}
