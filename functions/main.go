package main

import "fmt"

func main() {
	doSomething()
}

func doSomething() {
	fmt.Println("Do something")
	sum := sum(1, 2)
	fmt.Println(sum)
	fmt.Println(sumAll(1, 2, 3, 4, 5))
	sumAll, length := sumAndLength(10, 20, 30)
	fmt.Println("the sum is : ", sumAll)
	fmt.Println("the length is : ", length)
}

// function with input parameters
func sum(value1 int, value2 int) int {
	return value1 + value2
}

// If the parameters of the function are all same type, type needs be mentioned only once
// func sum(value1, value2 int) int {
// 	return value1 + value2
// }

// Get arbitary number of parameters
func sumAll(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

// A function can returns multiple values
func sumAndLength(values ...int) (int, int) {
	total := 0
	for _, value := range values {
		total += value
	}
	return total, len(values)
}
