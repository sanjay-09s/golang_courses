package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	var (
		tobe   bool       = false
		maxInt uint       = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)

	fmt.Printf("Type : %T  and value : %v\n", tobe, tobe)
	fmt.Printf("Type : %T  and value : %v\n", maxInt, maxInt)
	fmt.Printf("Type : %T  and value : %v\n", z, z)

	var s string

	fmt.Printf("%q", s)

}
