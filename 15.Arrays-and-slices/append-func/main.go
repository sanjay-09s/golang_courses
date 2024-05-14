package main

import "fmt"

func main() {

	s := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println("-----------")
	s = append(s, 7, 8, 9)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println("-----------")

	t := []interface{}{1, 2, 4.35, "Sanju", 't'}
	fmt.Println(t)
	fmt.Println(len(t))
	fmt.Println("----------")

	t = append(t, s)
	fmt.Println(t)
	fmt.Println(len(t))
	fmt.Println("---------")

}
