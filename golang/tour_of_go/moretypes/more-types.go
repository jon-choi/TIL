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
	jon_str := "jon"
	jon = &jon_str
	fmt.Println(*jon)     // "jon"
	fmt.Println(jon)      // 0xc42000e2b0
	fmt.Println(jon_str)  // "jon"
	fmt.Println(&jon_str) // 0xc42000e2b0
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

//
func arrays() {
}

func main() {
	// pointers()
	structs()
}
