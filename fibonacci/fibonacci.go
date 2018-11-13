package main

import "fmt"

func fibonacciMemoized() func() int {
	n0 := 0
	n1 := 1
	return func() int {
		n0, n1 = n1, n0+n1
		return n0
	}
}

func fibonacciNaive(n int) int {
	if n <= 2 {
		return 1
	}

	return fibonacciNaive(n-1) + fibonacciNaive(n-2)
}

func fibonacci(n int) int {
	n0 := 0
	n1 := 1

	for i := n; i > 0; i-- {
		n0, n1 = n1, n0+n1
	}

	return n0
}

func main() {
	f := fibonacciMemoized()
	for i := 0; i < 10; i++ {
		// fmt.Println(f())
		f()
	}

	fmt.Println(fibonacciNaive(10))
	fmt.Println(fibonacci(10))
}
