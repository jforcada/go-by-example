/* variables.go */

package main

import "fmt"

func main() {

	// var declares 1 or more variables.
	var a string = "initial"
	fmt.Println(a)

	// you can declare one or more variables at once.
	var b, c int = 1, 2
	fmt.Println(b, c)

	// go will infer the type of initialized variables
	var d = true
	fmt.Println(d)

	// Variables declared without a corresponding initialization are zero-valued.
	var e int
	fmt.Println(e)

	// The := syntax is shorthand for declaring and initializing a variable
	f := "short"
	fmt.Println(f)
}
