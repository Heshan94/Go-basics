// Maps are unordered data structure
package main

import (
	"fmt"
	"sort"
)

func main() {
	countries := make(map[string]string)
	fmt.Println(countries) // map[]

	countries["sl"] = "Sri Lanka"
	countries["aus"] = "Australia"
	countries["nz"] = "New Zeland"
	fmt.Println(countries) // map[aus:Australia nz:New Zeland sl:Sri Lanka]

	// Delete an element
	delete(countries, "sl")
	fmt.Println(countries) // map[aus:Australia nz:New Zeland]

	// Add a new element
	countries["ca"] = "Canada"
	fmt.Println(countries) // map[aus:Australia ca:Canada nz:New Zeland]

	// Iterate through the map
	for k, v := range countries {
		fmt.Printf("%v, %v\n", v, k)
	}

	// Stotre keys of the map in a slice and sort
	keys := make([]string, len(countries))
	i := 0
	for k := range countries {
		keys[i] = k
		i++
	}
	fmt.Println(keys) // [ca aus nz]
	sort.Strings(keys)
	fmt.Println(keys) // [aus ca nz]

	// Print values of the map in alphabatical order
	for i := range keys {
		fmt.Println(countries[keys[i]])
	}
}
