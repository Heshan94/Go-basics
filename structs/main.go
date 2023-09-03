// Structs in go is a custom type which similar to classes in other oop languages
// Structs can have both data and methods
// If the struct name starts with an uppercase character it means it's exported and other parts of the application can use it
// Lowercase means it's not accessible. This is similar to access modifiers in other oop languages
// This principle is valid for variable names and method names as well

package main

import (
	"fmt"
)

type Dog struct {
	Breed  string
	Weight int
}

func main() {
	poodle := Dog{"poodle", 10} // This is like the constructor of a class. Initlizing values will be given in the same order that they've been declared
	fmt.Println(poodle)         // {poodle 10}
	fmt.Printf("%+v\n", poodle) // {Breed:poodle Weight:10} : %+v will shows the variable names in the struct

	// Access fields in the struct
	fmt.Printf("Breeding :%v Weight : %v\n", poodle.Breed, poodle.Weight) // Breeding :poodle Weight : 10

	// Change data fields of the struct
	poodle.Weight = 20
	fmt.Printf("Breeding :%v Weight : %v\n", poodle.Breed, poodle.Weight) // Breeding :poodle Weight : 20
}
