package main

import "fmt"

func main() {
	m := make(map[string][]string)
	m["bond_james"] = []string{`shaken`, `not_stirred`, `fast cars`}
	m["moneypenny_jenny"] = []string{`a`, `b`, `c`}
	m[`Sanjay`] = []string{`sugan`, `sengodan`, `mani`}
	m[`sugan`] = []string{`hi`, `hello-8`}

	fmt.Printf("%#v\n", m)
	for k, v := range m {
		fmt.Println(k)
		for i, j := range v {
			fmt.Println(i, j)
		}
	}

	fmt.Println("-----------------------")
	delete(m, "Sanjay")
	fmt.Println("-----------------------")
	for k, v := range m {
		fmt.Println(k)
		for i, j := range v {
			fmt.Println(i, j)
		}
	}
}
