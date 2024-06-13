package main

import "fmt"

func main() {
	defer foo()
	bar()
	fmt.Println(`foo bar`)
}

func foo() {
	fmt.Println(`foo`)
}

func bar() {
	fmt.Println(`bar`)
}
