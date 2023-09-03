package main

import (
	"fmt"
)

func main() {
	x := 10
	var pointer *int = &x
	// pointer := &x This also works

	fmt.Println(x)
	fmt.Println(pointer)  // memory address
	fmt.Println(*pointer) // Dereferencing

	fNumber := 14.56
	fPointer := &fNumber
	fmt.Println(*fPointer)
}
