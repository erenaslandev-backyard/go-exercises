package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	checkError(scanner.Err())

	nTemp, err := strconv.ParseInt(scanner.Text(), 10, 64)
	checkError(err)
	n := int(nTemp)

	fmt.Println(convertToBase(n, 2))
}

func convertToBase(n int, base int) string {
	var resultBuilder strings.Builder
	var remainder int
	quotient := n

	for quotient > 0 {
		quotient, remainder = divmod(quotient, base)
		if remainder == 0 {
			resultBuilder.WriteRune('0')
		} else {
			resultBuilder.WriteString(strconv.Itoa(remainder))
		}
	}
	return reverse(resultBuilder.String())
}

func divmod(numerator int, denominator int) (quotient int, remainder int) {
	return numerator / denominator, numerator % denominator
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
