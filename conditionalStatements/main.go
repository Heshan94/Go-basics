// No parenthesis are wrapping the boolean check
package main

import "fmt"

func main() {
	x := 10
	if x < 0 {
		fmt.Println("Less than zero")
	} else if x == 0 {
		fmt.Println("Equal to zero")
	} else {
		fmt.Println("Greater than zero")
	}

	// A variable can be intialized in the if condition
	if x := -10; x < 0 {
		fmt.Println("Less than zero")
	} else if x == 0 {
		fmt.Println("Equal to zero")
	} else {
		fmt.Println("Greater than zero")
	}

}
