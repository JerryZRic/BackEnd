//current package
package main

//import packages
import (
	"fmt"
	"math"
	"math/rand"
)

//golang's basic types
var (
	aBool bool = true

	aString string = "something"

	int1 int
	int2 int8
	int3 int16
	int4 int64

	uint1 uint
	uint2 uint8 = 3
	uint3 uint16
	uint4 uint64
	aUintptr uintptr
	
	aByte byte // alias for uint8

	aRune rune // alias for int32, and a Unicode code point?

	//need a explicit conversion.
	float1 float32 = float32(uint2)
	float2 float64

	complex1 complex64
	complex2 complex128 = 24 + 16i
)

var aVarialbe, bVariable int
var cVarialbe, dVariable int = 3, 4

const (
	MyPi1, MyPi2 = 3.14, 3.14159
	BigNumber = 1 << 100
	SmallNumber = BigNumber >> 99
)

func main() {

	//as a pseudo random number, it always returns the same value.
	fmt.Println("Output a random number:", rand.Intn(100))

	fmt.Println("Output a square root:", math.Sqrt(9))

	fmt.Println("Output Pi start with a capital as an exported name:", math.Pi)

	//short variable declarations can be used inside a function, type is inferred from return values.
	addResultInt, addResultStr := myAdd(5, 99)
	fmt.Println("Call my add function with 2 parameters and 2 return values:", addResultInt, addResultStr)
	
	fmt.Println("Output package level variables (defaults to their zero value):", aVarialbe, bVariable)
	
	fmt.Println("Output package level variables with initializers:", cVarialbe, dVariable)

	fmt.Printf("Type: %T Value: %v\n", aBool, aBool)
	fmt.Printf("Type: %T Value: %v\n", aString, aString)
	fmt.Printf("Type: %T Value: %v\n", uint2, uint2)
	fmt.Printf("Type: %T Value: %v\n", float1, float1)
	fmt.Printf("Type: %T Value: %v\n", complex2, complex2)

	fmt.Println("Output constants:", MyPi1, MyPi2)

	needInt32 := func (n int32) int32 { return n } (SmallNumber)
	needFloat64 := func (n float64) float64 { return n } (SmallNumber)
	fmt.Printf("untyped  numberic constants takes the type needed its context: %T -- %v, %T -- %v\n", needInt32, needInt32, needFloat64, needFloat64)
}

//x and y share a type int.
//named return values.
func myAdd (x, y int) (valueInt int, valueString string) {
	valueInt = x + y
	valueString = "multiple results"

	//this function need a naked return.
	return
}