// The break statment is nor required in each case statement
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix()) // This line of code seeds the random number generator with the current Unix timestamp to ensure random number generation varies across program execution
	dow := rand.Intn(7) + 1
	fmt.Println(dow)
	var result string

	switch dow {
	case 1:
		result = "It's Sunday"
	case 2:
		result = "It's Monday"
	default:
		result = "It's any other day"
	}
	fmt.Println(result)
}
