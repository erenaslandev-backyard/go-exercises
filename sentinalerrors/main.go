package main

import "fmt"

type Error string

func (e Error) Error() string {
	return string(e)
}

const anErr = Error("I AM AN ERROR")

func main() {
	// Go constants are COOL!
	fmt.Println(anErr)
}
