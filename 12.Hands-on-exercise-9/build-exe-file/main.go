package main

import "fmt"

func main() {
	fmt.Println("Hello Gophers. This is from terminal")
}

// add the exe file using the build "go build main.go"
// to run the exe file "./main.exe"
// to remove the file rm "rm main.exe"
// "GOOS = windows go build ." create the file in windows, does not create in mac or linux
// "GOOS = darwin go build ." create the file in mac, does not create in linux or windows
// "GOOS = linux go build ." create the file in linux, does not create in mac or windows
