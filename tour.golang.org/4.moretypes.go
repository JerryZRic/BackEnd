package main

import (
	"fmt"
	"strings"
	"strconv"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	var a  = 333
	p := &a
	*p = 222
	fmt.Println("pointer p's underlying value:", *p)
	fmt.Println("a's value:", a)

	//X:0 is implicit
	v := Vertex{Y: 555}
	vp := &v
	fmt.Println("sturct", *vp)
	fmt.Println("vertex x:", vp.X, "vertex Y:", vp.Y)

	var aArray [2]int
	aArray[0] = 5
	aArray[1] = 7
	bArray := []int{1, 2, 3, 4, 5, 6, 7}
	cArray := bArray[2:4]
	fmt.Println("aArray:", aArray)
	fmt.Printf("bArray:%v, length:%v, capacity:%v\n", bArray, len(bArray), cap(bArray))
	bArrayAppend := append(bArray, 99, 88, 77, 66, 55, 44, 33, 22)
	fmt.Printf("bArray append item: %v, length:%v, capacity:%v\n", bArrayAppend, len(bArrayAppend), cap(bArrayAppend))
	fmt.Println("bArray[2:4]:", cArray)
	fmt.Printf("bArray[2:4] length:%v, capacity:%v\n", len(cArray), cap(cArray))
	bArray[2] = 222
	fmt.Println("bArray[2:4] after bArray changed:", cArray)
	dArray := bArray[:]
	eArray := bArray[2:]
	fArray := bArray[:4]
	fmt.Println("bArray[:]:", dArray)
	fmt.Println("bArray[2:]:", eArray)
	fmt.Println("bArray[:4]:", fArray)

	var nilSlice []int
	fmt.Printf("nil slice length:%v, capacity:%v\n", len(nilSlice), cap(nilSlice))

	makeSlice := make([]int, 2, 5)
	fullArray := makeSlice[:cap(makeSlice)]
	fmt.Printf("makeSlice: %v length:%v, capacity:%v\n", makeSlice, len(makeSlice), cap(makeSlice))
	fmt.Printf("fullArray: %v length:%v, capacity:%v\n", fullArray, len(fullArray), cap(fullArray))

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "0"
	board[0][2] = "X"
	for i := 0; i < len(board); i++ {
		fmt.Printf("%v\n", strings.Join(board[i], " "))
	}

	for i, v := range bArray {
		fmt.Printf("each iteration index:%v, value:%v\n", i, v)
	}
	for i, _ := range bArray {
		fmt.Printf("each iteration index:%v\n", i)
	}
	for _, v := range bArray {
		fmt.Printf("each iteration value:%v\n", v)
	}

	aMap := make(map[string] int)
	aMap["a"] = 12
	aMap["b"] = 55
	fmt.Println("aMap:", aMap)

	bMap := map[string] Vertex{
		"a": {34, 3},
		"b": {54, 2},
	}
	fmt.Println("bMap:", bMap)

	bMapValue, hasKey := bMap["b"]
	fmt.Println("bMap has key b?", hasKey, "value:", bMapValue)
	delete(bMap, "b")
	bMapValue, hasKey = bMap["b"]
	fmt.Println("bMap has key b?", hasKey, "value:", bMapValue)

	fmt.Println(callFunc(funcAdd))
}

func callFunc(aFunc func(int, int) int) string {
	return "Call function result: " + strconv.Itoa(aFunc(1, 2))
}

func funcAdd(a, b int) int {
	return a + b
}