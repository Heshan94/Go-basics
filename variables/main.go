package main

import "fmt"

const myConst int = 100

func main() {
	var myStr string = "hello"
	fmt.Println(myStr)
	fmt.Printf("Type of myStr : %T\n", myStr)

	var myStr2 = "hello 2" // The compiler will determine the type
	fmt.Printf("Type : %T\n", myStr2)

	mystr3 := "Another hello" // This can be only used inside a function
	fmt.Printf("%T\n", mystr3)

	var x int = 10
	fmt.Println(x)

	fmt.Println(myConst)

}
