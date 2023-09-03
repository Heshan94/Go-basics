package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// When we need to read user inputs from the command line all the values will be considered as strings
// We need to convert the string to other types using strconv and strings library

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a number: ")
	input, _ := reader.ReadString('\n')

	// ParseFloat() in strconv is used to convert the string into a float
	// It takes two parameters: 1) the string 2) Number of bytes of the float : 32 or 64
	// strings.TrimSpace() is used to remove the spaces in the string
	// strconv.ParseFloat() returns an error, err variable will have the value and it won't be nill
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(aFloat)
	}
}
