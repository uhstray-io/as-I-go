package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	zSq    complex128 = cmplx.Sqrt(-5 + 12i)
)

func swap(x, y string) (string, string) {
	return y, x
}

func add(x, y int) int {
	return x + y
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func helloworld() {
	fmt.Println("\n===HELLO WORLD===")

	fmt.Println("Hello World!")
	fmt.Println("The time is", time.Now())
}

func earlymaths() {
	fmt.Println("\n===EARLY MATHS===")

	var x = rand.Intn(111)
	fmt.Println("A random number is", x)

	var y = math.Sqrt(float64(x))
	fmt.Println("Now you have %g problems.\nSQRT:", y)
	fmt.Println("PI:", math.Pi)

	var z = add(x, int(y))
	fmt.Println("ADD:", z)

	m, n := split(z)
	fmt.Println("SPLIT:", m, n)
}

func strings() {
	fmt.Println("\n===STRINGS===")

	a, b := swap("hello", "world")
	fmt.Println("SWAP:", a, b)
}

func variables() {
	// BASIC GO TYPES:
	// bool
	// string (utf-8)
	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte // alias for uint8
	// rune // alias for int32 // doubles as string (utf-8)
	// float32 float64
	// complex64 complex128

	var i, j int = 1, 2
	k := 3
	var c, python bool = true, false
	var java string = "no!"

	fmt.Println("\n===VARIABLES===")

	fmt.Println("INT:", i, j, k, "\nBOOL:", c, python, java)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", zSq, zSq)
}

func zerovars() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Println("\n===ZERO'D OUT VARIABLES===")
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

func main() {
	helloworld()
	earlymaths()
	strings()
	variables()
	zerovars()
}
