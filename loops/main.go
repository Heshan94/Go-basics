// Go supports various for loops including the traditional three statement one.
// However, there is separate loop called while in go. It uses the same for loop to achieve it.
// break and continue statements work the same
package main

import (
	"fmt"
)

func main() {
	colors := []string{"red", "blue", "green"}

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	// range based for loop
	for i := range colors {
		fmt.Println(colors[i])
	}

	// for-each loop
	// This returns two values. 1) index 2) value
	// _ is isued ignore the index part
	for _, color := range colors {
		fmt.Println(color)
	}

	// Simliar to while loop
	count := 0
	for count < 10 {
		fmt.Println(count)
		count++
	}

	// Both break and continue statements work the same purpose as other languages
	// Additonaly there is goto statement which will break the loop and go to some other part of the code

	sum := 1

	for sum < 1000 {
		sum++
		sum += sum
		fmt.Println(sum)
		if sum > 200 {
			goto toEnd
		}
	}
toEnd:
	fmt.Println("This is the end")
}
