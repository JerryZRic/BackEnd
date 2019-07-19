package main

import (
	"fmt"
	"math/rand"
	"math"
	"runtime"
)

func main() {

	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
	}
	fmt.Println("run a for loop to add from 1 to 100:", sum)

	sum = 1
	for ; sum < 1000; {
		sum += 1
	}
	fmt.Println("run a for loop omit init and post statements:", sum)

	sum = 10
	for sum < 1000 {
		sum += rand.Intn(100)
	}
	fmt.Println("run go's \"while\" loop:", sum)

	sum = 22
	for {
		//if with a short statement
		if n := math.Pow(33, 2); sum > int(n) {
			break
		} else {
			sum += rand.Intn(9)
		}
	}
	fmt.Println("run a forever loop with break:", sum)

	fmt.Print("runing platform is: ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("MAC OS.")
	case "linux":
		fmt.Println("Linux.")
	case "windows":
		fmt.Println("Microsoft Windows.")
	default:
		fmt.Printf("%s.\n", os)
	}

	//evaluate cases from top to bottom.
	switch n := 4; n {
	case func() int { fmt.Println("Run case 1."); return 1 }():
		fmt.Println("Hit 1.")
	case func() int { fmt.Println("Run case 2."); return 2 }():
		fmt.Println("Hit 2.")
	case func() int { fmt.Println("Run case 3."); return 3 }():
		fmt.Println("Hit 3.")
	case func() int { fmt.Println("Run case 4."); return 4 }():
		fmt.Println("Hit 4.")
	case func() int { fmt.Println("Run case 5."); return 5 }():
		fmt.Println("Hit 5.")
	}

	//no condition switch as long if-then-else chain
	switch n := 19; {
	case n < 10:
		fmt.Println("n < 10")
	case n < 20:
		fmt.Println("10 < n < 20")
	case n < 30:
		fmt.Println("20 < n < 30")
	case n < 40:
		fmt.Println("30 < n < 40")
	default:
		fmt.Println("n > 40")
	}

	aInt := 123;
	//param aInt is evaluated immediately
	defer fmt.Println("FIRST call pushed onto stack")
	defer fmt.Println("aInt in defer is:", aInt)
	defer fmt.Println("LAST call pushed onto stack")
	aInt = 456;
	fmt.Println("aInt is:", aInt)
	aInt = 789;
	fmt.Println("aInt is:", aInt)
}