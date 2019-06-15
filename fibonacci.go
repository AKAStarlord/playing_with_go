package main

import (
	"fmt"
	"time"
)

func main() {
	FIB_NUM := 15

	// Start a timer and run Fib alg 2x sequentially and report time.
	startTime0 := time.Now()
	fmt.Println(FibonacciRecursion(FIB_NUM))
	fmt.Println(FibonacciRecursion(FIB_NUM))
	endTime0 := time.Now()
	fmt.Println(endTime0.Sub(startTime0))

	// Start a timer and run Fib alg 2x *in parallel* and report time.
	startTime1 := time.Now()
	go fmt.Println(FibonacciRecursion(FIB_NUM))
	go fmt.Println(FibonacciRecursion(FIB_NUM))
	endTime1 := time.Now()
	fmt.Println(endTime1.Sub(startTime1))
}

func FibonacciRecursion(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
}
