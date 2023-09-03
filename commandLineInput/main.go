// To do this, we need Reader object from bufio and os to get what user types in
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	// ReadString() will return two values. input and error. If you want to ignore the second value use '_'
	//  '/n' measn it will read untill user enters Enter
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
}
