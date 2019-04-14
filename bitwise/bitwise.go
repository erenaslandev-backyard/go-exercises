package main

import "fmt"

const maxInt = int(^uint(0) >> 1)
const minInt = -maxInt - 1

func main() {
	fmt.Printf("MaxInt: %d\n", maxInt)
	fmt.Printf("MinInt: %d\n", minInt)

	fmt.Println(uint(0))
	fmt.Println(^uint(0))
	fmt.Println(^uint(0) >> 1)
	fmt.Println(^0)
	fmt.Println(^1)
	fmt.Println(^2)
	fmt.Println(0 >> 1)
	fmt.Println(1 >> 2)
}
