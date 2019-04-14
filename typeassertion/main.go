package main

import "fmt"

type helloer interface {
	sayHello()
}

type myString string

func (s myString) sayHello() {
	fmt.Println("Hello:", s)
}

func explain(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("type string")
	case myString:
		fmt.Println("type myString")
	}
}

func main() {
	hello := "hello world"
	myHello := myString(hello)
	// hello.(string) compiler error: not an interface type
	fmt.Println(string(hello))
	// int(hello) compiler error: cannot convert
	// fmt.Println(helloer(hello).(helloer))
	fmt.Println(helloer(myHello).(helloer))
	val, ok := helloer(myHello).(helloer)
	if ok {
		fmt.Println(val)
	}

	fmt.Println(hello)
	// hello.sayHello()

	explain(hello)
	explain(myHello)
}
