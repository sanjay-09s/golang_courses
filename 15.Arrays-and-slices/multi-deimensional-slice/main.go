package main

import "fmt"

func main() {
	jb := []string{"Sanjay", "Sugan", "Sengodan", "Mani"}
	fmt.Println(jb)
	fmt.Println("-----------")

	jm := []string{"Sanju", "Sugu"}
	fmt.Println(jm)
	fmt.Println("-----------")

	dm := [][]string{jm, jb}
	fmt.Println(dm)
	fmt.Println("-----------")

	db := [][]string{{"A", "B"}, {"C", "D"}}
	fmt.Println(db)
	fmt.Println("-----------")
}
