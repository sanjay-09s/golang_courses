package main

import "fmt"

func main() {
	adams := 42
	fmt.Printf("42 as binary, %b \n", adams)
	fmt.Printf("42 as hexadecimal, %x \n", adams)

	//Print these values as both binary & hexadecimal
	a, b, c, d := 0, 1, 2, 30
	fmt.Printf("%v \t %b \t %X \n", b, b, b)
	fmt.Printf("%v \t %b \t %X \n", c, c, c)
	fmt.Printf("%v \t %b \t %X \n", a, a, a)
	fmt.Printf("%v \t %b \t %X \n", d, d, d)

	// inorder to give with hexadecimal values
	fmt.Println("")
	a, b, c, d = 0, 1, 2, 230
	fmt.Printf("%v \t %b \t %#X \n", b, b, b)
	fmt.Printf("%v \t %b \t %#X \n", c, c, c)
	fmt.Printf("%v \t %b \t %#X \n", a, a, a)
	fmt.Printf("%v \t %b \t %#X \n", d, d, d)

}
