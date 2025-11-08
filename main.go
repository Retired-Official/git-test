package main

import (
	"fmt"
	"git-test/src"
)

func main() {
	name := "Michelle"
	fmt.Printf("Before reverse: %v\n", name)
	reversedName := src.Reverse(name)
	fmt.Printf("After reverse: %v\n", reversedName)
}