/* for */

package main

import "fmt"

func main() {

	// Single condition
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// Classic initial/condition/after
	for j:= 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// Without condition. Will loop until break or return from the
	// closing function.
	for {
		fmt.Println("loop")
		break
	}
}
