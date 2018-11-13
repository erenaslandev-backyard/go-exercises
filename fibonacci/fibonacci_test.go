package main

import (
	"fmt"
	"testing"
)

func ExampleFibonacci() {
	fmt.Println(fibonacci(10))
	// Output
	// 55
}
func TestFibonacci(t *testing.T) {
	if fibonacci(10) != 55 {
		t.Error("You are wrong!")
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fibonacci(10)
	}
}

func ExampleFibonacciNaive() {
	fmt.Println(fibonacciNaive(10))
	// Output
	// 55
}

func TestFibonacciNaive(t *testing.T) {
	if fibonacciNaive(10) != 55 {
		t.Error("You are wrong!")
	}
}
func BenchmarkFibonacciNaive(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fibonacciNaive(10)
	}
}
