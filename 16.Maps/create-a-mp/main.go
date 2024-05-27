package main

import "fmt"

func main() {
	am := map[string]int{
		"sanjay": 21,
		"sugan":  19,
		"Sanju":  23,
	}
	fmt.Println(am)
	fmt.Printf("%#v\n", am)
	fmt.Println(len(am))

	an := make(map[string]int)
	an["mani"] = 40
	an["Daddy"] = 55
	fmt.Println(an)
	fmt.Printf("%#v\n", an)
	fmt.Println(len(an))
}
