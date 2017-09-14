package main

import "fmt"

// A pointer holds he memeory address of a value.
func pointers() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // point to print i through the pointer
	fmt.Println(p)  // print the memory location
	*p = 21         // set i throught the pointer
	fmt.Println(i)  // set the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide through the pointer
	fmt.Println(j) // see the new value of j

	var jon *string
	jonStr := "jon"
	jon = &jonStr
	fmt.Println(*jon)    // "jon"
	fmt.Println(jon)     // 0xc42000e2b0
	fmt.Println(jonStr)  // "jon"
	fmt.Println(&jonStr) // 0xc42000e2b0
}

func structs() {
	// A struct is a collection of fields.
	type Vertex struct {
		X int
		Y int
	}

	v := Vertex{1, 2}
	p := &v // accessing structs through a pointer
	p.X = 1e9
	fmt.Println(v)
	fmt.Println(v.X)
	fmt.Println(p.X)    // no need for explicit (*p)
	fmt.Println((*p).X) // need the parens.

	var (
		v1      = Vertex{1, 2}
		v2      = Vertex{X: 1}  // Y:0 is implicit
		v3      = Vertex{}      // X:0 and Y:0
		pointer = &Vertex{1, 2} // has type *Vertex
	)

	fmt.Println(v1, pointer, v2, v3)
}

// The type [n]T is an array of n values of type T.
func arrays() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

// Unlike arrays which have fixed sizes. A slice is dynamically-sized.
func slices() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int
	s = primes[1:4]
	fmt.Println(s)
}

func slicesAreRefs() {
	names := [4]string{
		"Eugene",
		"Tae",
		"P",
		"I",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	c := names[0:4]
	fmt.Println(a, b)
	fmt.Println(c)

	b[0] = "XXX" // This changes names[1]
	fmt.Println(a, b)
	fmt.Println(names)
}

// Array literal without the length.
// Creates an equivalent array and a slice that refers to it.
func sliceLiterals() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, false}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

func sliceDefaults() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// Length is number of elements the slice contains.
// Capacity is number of elements in underlying array.
func sliceLengthAndCapacity() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:0]
	printSlice(s)

	// You can exten a slice's length by re-slicing it, provided it has sufficient capacity.
	s = s[:4] // len 4. cap 6
	printSlice(s)

	s = s[2:] // len 2. cap 4
	printSlice(s)
}

func main() {
	// pointers()
	// structs()
	// arrays()
	// slices()
	// slicesAreRefs()
	// sliceLiterals()
	// sliceDefaults()
	sliceLengthAndCapacity()
}
