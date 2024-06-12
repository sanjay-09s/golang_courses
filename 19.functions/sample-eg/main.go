package main

import "fmt"

func main() {
	foo()
	bar("Shinchan Yamata")
	s := aloha("Shinchan")
	fmt.Println(s)
	s1, u := dogYears("Shiro", 3)
	fmt.Println(s1, u)
}

func foo() {
	fmt.Println("I am from Foo()")
}

func bar(s string) {
	fmt.Println("My name is ", s)
}

func aloha(s string) string {
	return fmt.Sprint("They call me Mr.", s)
}

func dogYears(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprint(name, " is dogname which is old of years"), age
}

/*

func (r receiver) (identifier(p parameters)) (return type){ <your_code> }

no params, no return type
1 params, no return type
1 params, 1 return type
2 params, 2 return type
*/
