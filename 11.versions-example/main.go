package main

import "fmt"

func main() {
	fmt.Println("Hello, we're working on vaersion tags")
}

// "git tag" shows nothing when there is no tag
// To set up the tag "git tag v1.0.0" here we can use any version
// to see info about the tag "git show v1.0.0"

// also we can get the latest module dependices using the commit ID
// eg: "go get github.com/uni-SanjayS/puppy@commit_ID"

// To get the latest commit "go get github.com/uni-SanjayS/puppy@latest"
// to get the particular version of that code "go get github.com/uni-SanjayS/puppy@v1.0.0"
