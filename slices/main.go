// Slices are resizeable and it's possible to add, remove and sort elements

package main

import (
	"fmt"
	"sort"
)

func main() {
	// declaring a slice is similar to array but we don't pass the size inside []
	colors := []string{"red", "blue", "green"}
	fmt.Println(colors) // [red blue green]

	// add an element
	colors = append(colors, "white")
	fmt.Println(colors) // [red blue green white]

	// remove the first element
	colors = append(colors[1:len(colors)])
	fmt.Println(colors) // [blue green white]

	// remove the last element
	colors = append(colors[:len(colors)-1])
	fmt.Println(colors) // [blue green]

	// The build in make() function can be used create slices. It takes 3 parameters.
	// 1) type of the slice 2) initial size 3) (optional) max size

	numbers := make([]int, 3, 5)
	fmt.Println(numbers) // [0, 0, 0]

	numbers[0] = 3
	numbers[1] = 1
	numbers[2] = 2
	fmt.Println(numbers) // [3 1 2]

	// Sort a slice
	sort.Ints(numbers)
	fmt.Println(numbers) // [1 2 3]
}
