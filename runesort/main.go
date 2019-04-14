package main

import (
	"fmt"
	"sort"
)

type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

var runes RuneSlice = []rune{'可', '愛'} // sort based on the integer values of the runes

func main() {
	// Runes
	fmt.Println("Original : ", runes[:])
	fmt.Println("Original : ", string(runes[:]))

	sort.Sort(RuneSlice(runes))

	fmt.Println("Sort : ", runes[:])
	fmt.Println("Sort : ", string(runes[:]))

	sort.Sort(sort.Reverse(runes[:]))

	fmt.Println("Reverse : ", runes[:])
	fmt.Println("Reverse : ", string(runes[:]))

	for i, v := range "h可llo" {
		fmt.Printf("%d %T %v %c\n", i, v, v, v)
	}
}
