package main

import "fmt"

func main() {

	xi := []int{2, 4, 5, 32, 1, 24, 14, 1, 41, 4}

	for i, j := range xi {
		fmt.Printf("Index --> %v : Value --> %v\n", i, j)
	}

	fmt.Println(xi[5])

	v := map[string]int{
		"James": 34,
		"Diana": 28,
	}

	for i, j := range v {
		fmt.Printf("Key --> %v : Value --> %v\n", i, j)
	}

	// will work when the value is present and set the "ok" --> true
	if m, ok := v["James"]; ok {
		fmt.Println("The value with the key \"JAMES\" is present, the value", m)
	}

	// will work when the value is not present and set the "ok" --> false
	if m, ok := v["L"]; !ok { // "!ok" retunns true
		fmt.Println("There is no value in the map with \"L\" and the zero value is", m)
	}

}
