package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("s - %#v\n", s)
	fmt.Println(len(s))
	fmt.Println("-----------")
	s = append(s, 7, 8, 9)
	fmt.Printf("s - %#v\n", s)
	fmt.Println(len(s))
	fmt.Println("-----------")

	t := []interface{}{1, 2, 4.35, "Sanju", 't'}
	fmt.Printf("t - %#v\n", t)
	fmt.Println(len(t))
	fmt.Println("----------")

	t = append(t, s)
	fmt.Printf("t - %#v\n", t)
	fmt.Println(len(t))
	fmt.Println("---------")

}

// output => []interface {}{1, 2, 4.35, "Sanju", 116, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}}
