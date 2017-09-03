package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func forLoop() {
	sum := 1
	// init, condition, post
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func while() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func infiniteLoop() {
	for {
	}
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func newtonSqrtBaseFormula(x float64, start float64) float64 {
	return start - (start*start-x)/(2*start)
}

func NewtonSqrtEst(x float64) float64 {
	z := 20.0
	// const limit = 10
	// for i := 0; i < limit; i += 1 {
	// 	z = NewtonSqrt(x, z)
	// }
	iter := 0
	fmt.Println(z - newtonSqrtBaseFormula(x, z))
	for z-newtonSqrtBaseFormula(x, z) > 0.00001 {
		z = newtonSqrtBaseFormula(x, z)
		iter += 1
		fmt.Printf("%d iterations\n", iter)
	}
	fmt.Printf("%d iterations\n", iter)
	return z
}

func osSwitch() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows ...
		fmt.Printf("%s.", os)
	}
}

func sunday() {
	fmt.Println("when's Sunday?")
	today := time.Now().Weekday()
	switch time.Sunday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Far away...")
	}
}

func timeOfDay() {
	t := time.Now()
	hour := t.Hour()
	switch {
	case hour < 12:
		fmt.Println("Good Morning!")
	case hour < 17:
		fmt.Println("Good Afternoon!")
	default:
		fmt.Println("Good Evening!")
	}
}

func helloWhat() {
	defer fmt.Println("world")
	fmt.Println("hello")
}

func countDefer() {
	fmt.Println("Counting...")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Done")
}

func main() {
	// forLoop()
	// while()
	// infiniteLoop()
	// fmt.Println(sqrt(2), sqrt(-4))
	// fmt.Println(
	// 	pow(3, 2, 10),
	// 	pow(3, 3, 20),
	// )
	// fmt.Printf("%f is estimate", NewtonSqrtEst(2))
	// osSwitch()
	// sunday()
	// timeOfDay()
	// helloWhat()
	countDefer()
}
