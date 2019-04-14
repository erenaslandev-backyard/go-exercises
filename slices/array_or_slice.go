package main

import "fmt"
import "reflect"

func main() {
	m := make(map[string]interface{})
	m["a"] = []string{"a", "b", "c"}
	m["b"] = [4]int{1, 2, 3, 4}
	m["c"] = []int{1, 2, 3, 4}
	m["d"] = make([]int, 4)
	m["e"] = new([]int)

	test(m)

	slice := []int{1, 2, 3, 4}
	testArgument(slice)

	array := [4]int{1, 2, 3, 4}
	// testArgument(array) // cannot use array (type [4]int)

	testVariadic(slice...)
	testVariadic(array[:]...)
}

func test(m map[string]interface{}) {
	for k, v := range m {
		rt := reflect.TypeOf(v)
		switch rt.Kind() {
		case reflect.Slice:
			fmt.Println(k, "is a slice with element type", rt.Elem())
		case reflect.Array:
			fmt.Println(k, "is an array with element type", rt.Elem())
		case reflect.Ptr:
			fmt.Println(k, "is a pointer with element type", rt.Elem())
		default:
			fmt.Println(k, "is something else entirely")
		}
	}
}

func testArgument(v []int) {
	rt := reflect.TypeOf(v)
	fmt.Println("SLICE ARGUMENT")
	fmt.Println("Value:", v)
	fmt.Println("Kind:", rt.Kind())
	fmt.Println("Type:", rt.Elem())
}

func testVariadic(v ...int) {
	rt := reflect.TypeOf(v)
	fmt.Println("VARIADIC")
	fmt.Println("Value:", v)
	fmt.Println("Kind:", rt.Kind())
	fmt.Println("Type:", rt.Elem())
}
