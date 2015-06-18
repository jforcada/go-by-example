package main

import "fmt"

func main() {

	// To create an empty map use: make(map[key-type]val-type)
	m := make(map[string]int)

	// Set key/value pairs with name[key] = val
	m["k1"] = 7
	m["k2"] = 13

	// Show all key value pairs
	fmt.Println("map:", m)

	// Get the value for a key
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// Number of key/value pairs
	fmt.Println("len:", len(m))

	// Built in delete removes key/value pairs
	delete(m, "k2")
	fmt.Println("map:", m)

	// Second argument is boolean, indicates if key was present.
	// To desambiguate between missing keys and keys with zero values.
	// Blank identifier _ to ignore the first value.
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// Declare and initialize a map in one line.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
