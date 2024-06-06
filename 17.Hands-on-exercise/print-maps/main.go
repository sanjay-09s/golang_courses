package main

import "fmt"

func main() {
	m := make(map[string][]string)
	m["bond_james"] = []string{`shaken`, `not_stirred`, `fast cars`}
	m["moneypenny_jenny"] = []string{`a`, `b`, `c`}
	m[`Sanjay`] = []string{`sugan`, `sengodan`, `mani`}

	fmt.Printf("%#v\n", m)
	for k, v := range m {
		fmt.Println(k)
		for i, j := range v {
			fmt.Println(i, j)
		}
	}
}
