package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// to convert the integer into the time
	// we have to use the time duration
	d1 := time.Duration(rand.Int63n(256))
	d2 := time.Duration(rand.Int63n(256))

	// creating the channels
	ch1 := make(chan int)
	ch2 := make(chan int)

	// runs in random time
	// go routines executes the code parallely
	go func() {
		time.Sleep(d1 * time.Millisecond)
		ch1 <- 41
	}()

	go func() {
		time.Sleep(d2 * time.Millisecond)
		ch2 <- 45
	}()

	// select statement will select which set of possible send or recieve operations will proceed.
	// It looks similar to the "switch" statement
	select {
	case v1 := <-ch1:
		fmt.Println("value from channel 1 ", v1)
	case v2 := <-ch2:
		fmt.Println("value from channel 2 ", v2)
	}

}
