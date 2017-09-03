package main

import (
	"fmt"
	// "math"
	"math/cmplx"
)

// var c, python, java bool

// var i, j int = 1, 2

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// func add(x int, y int) int {
func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - y
	return
}

func intToFloat(i int) (f float64) {
	f = float64(i)
	return
}

func needInt(x int) int           { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

func main() {
	// fmt.Println("My favorite number is", rand.Intn(100))

	// fmt.Printf("Now you have %g problems.", math.Sqrt(6))

	// fmt.Println(math.Pi)

	// fmt.Println(add(42, 13))

	// a, b := swap("hello", "world") // assignment & declaration
	// fmt.Println(a, b)

	// fmt.Println(split(17))

	// var i int
	// fmt.Println(i, c, python, java)

	// var c, python, java = true, false, "no!"
	// fmt.Println(i, j, c, python, java)

	// var i, j int = 1, 2
	// k := 3 // short assignment statement only available in function scope because statements have to begin with keywords
	// 			 // otherwise.
	// c, python, java := true, false, "no!"

	// fmt.Println(i, j, k, c, python, java)

	// fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	// fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	// fmt.Printf("Type: %T Value: %v\n", z, z)

	// i := 20
	// a := intToFloat(i)
	// fmt.Println(a)
	// fmt.Printf("%T", a)

	// 	const World = true
	// 	fmt.Println(World)
	const (
		Big   = 1 << 100
		Small = Big >> 99
	)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
