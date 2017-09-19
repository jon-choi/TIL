package main

import (
	"fmt"
	"math"
)

type vertex struct {
	X, Y float64
}

type myFloat float64

// Method is just a function with a receiver argument.
// Go does not have classes, but can define methods on types.
func (v vertex) absMethod() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func absFunction(v vertex) float64 {
	return v.absMethod()
}

// Can have the same name since different receiver argument.
// It has to refer to type defined in the same package.
func (f myFloat) absMethod() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Pointer receivers refer to the originally allocated value.
func (v *vertex) scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Value receivers refer to a copy.
func (v vertex) scaleValue(f float64) vertex {
	x := v.X * f
	y := v.Y * f
	return vertex{x, y}
}

// function
// Functions with pointer argument must take a pointer.
// Methods with pointer receivers can take either a value or a pointer.
// v.scaleValue(f) is called as (&v).scaleValue(f)
func scale(v *vertex, f float64) {
	v.X = f * v.X
	v.Y = f * v.Y
}

// Reasons to use a pointer receiver:
// #1 Method can modify the value that its receiver points to.
// #2 Avoid copying the value on each method call. More efficient for large structs.
// Tip: All methods on a given type hsould have either value or pointer receivers and don't mix.
func main() {
	// v := vertex{3, 4}
	// fmt.Println(v.absMethod())
	// fmt.Println(absFunction(v))

	// f := myFloat(-4.9)
	// fmt.Println(f.absMethod())

	v := vertex{3, 4}
	// fmt.Println(v)
	// v.scale(2)
	// fmt.Println(v)
	// newv := v.scaleValue(3)
	// fmt.Println(newv)

	// scale(&v, 10) // points to the reference.
	// fmt.Println(absFunction(v))
	// fmt.Printf("%v\n", v)
	// p := &v
	// fmt.Printf("%p\n", &v)
	// fmt.Printf("%v\n", p)
	// fmt.Printf("%p\n", &p)

	// fmt.Println(v.absMethod())  // works
	// fmt.Println(absFunction(v)) // works

	p := &vertex{4, 3}
	fmt.Println(p.absMethod()) // works
	fmt.Println(absFunction(*p))

	fmt.Println(*p)
	fmt.Println(p)
	fmt.Println(&p) // prints an address
	// fmt.Println(*v) // error
	fmt.Println(v)
	fmt.Println(&v) // why doesn't this?

	var t *vertex
	t = &v
	fmt.Printf("%p\n", t)  // 0x139839..   address of v
	fmt.Printf("%v\n", *t) // {3, 4}    value at address
	fmt.Printf("%v\n", &t) // 0x2133kj..  address of t (double pointer?)

	// Remember
	fmt.Println("*t denotes value at address of t")
	fmt.Println("&t address of t")
	fmt.Println("* ~ value at")
	fmt.Println("& ~ address of")
}
