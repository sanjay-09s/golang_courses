package main

import (
	"fmt"
	"log"
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

func logInfo(s fmt.Stringer) {
	log.Println("LOG FORM @20", s.String())
}

func main() {
	b := book{
		title: "Steal Like An Artist",
	}
	var c count = 65

	logInfo(b)
	logInfo(c)
}
