// Functions are declared outside the struct
// after the func keyword the receiver type needs to be passed. Here the Dog type
package main

import "fmt"

type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

func (d Dog) MakeSound() {
	fmt.Println(d.Sound)
}

func (d Dog) MakeSoundThreeTimes() {
	d.Sound = fmt.Sprintf("%v %v %v", d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}
func main() {
	dog := Dog{"poodle", 10, "wuf"}
	fmt.Println(dog)
	dog.MakeSound()
	dog.Sound = "yee"
	dog.MakeSound()
	dog.MakeSoundThreeTimes()
	dog.MakeSoundThreeTimes() // This won't bark 9 times since every time a copy of a Dog object is passed to the function.
	// If you need modify the actual object, then pointer type needs to be passed. Ex:
	// func (d *Dog) MakeSoundThreeTimes() {}

}
