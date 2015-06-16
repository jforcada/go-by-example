package main

import "fmt"

func main() {

	// Slices are typed only by the elements they contain, not the number
	// of elements). To create an empty slice with non-zero length, use
	// the builtin make.
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// We can set and get just like with arrays.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// len returns the length of the slice as expected.
	fmt.Println("len:", len(s))

	// In addition to these basic operations, slices support several more
	// that make them richer than arrays.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Slices can also be copy’d. Here we create an empty slice c of the
	// same length as s and copy into c from s.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Slices support a “slice” operator with the syntax slice[low:high].
	// Slice of the elements s[2], s[3], and s[4].
	l := s[2:5]
	fmt.Println("sl1:", l)

	// Slices up to, and excluding s[5]
	l = s[:5]
	fmt.Println("sl2:", l)

	// And slices up from, and including s[2]
	l = s[2:]
	fmt.Println("sl3:", l)

	// Slices can be composed into multi-dimensional data structures.
	// The length of the inner slices can vary, unlike with multi-dimensional
	// arrays.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
