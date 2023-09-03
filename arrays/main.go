package main

import "fmt"

func main() {
	var colors [3]string
	colors[0] = "red"
	colors[1] = "green"
	colors[2] = "blue"
	fmt.Println(colors) // [red green blue]

	// Initialize while declaring the array
	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers) // [1 2 3 4 5]

	// Length of an array
	length := len(numbers)
	fmt.Println(length) // 5
}
