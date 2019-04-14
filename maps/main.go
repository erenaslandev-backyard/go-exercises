package main

import "fmt"

func main() {
	set := make(map[string]struct{})
	for _, value := range []string{"apple", "orange", "apple"} {
		set[value] = struct{}{}
	}
	fmt.Println(set)
	// Output: map[orange:{} apple:{}]

	seen := make(map[string]struct{})
	for _, v := range []string{"apple", "orange", "apple"} {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
		}
	}

	for k := range seen {
		fmt.Printf("I saw %s\n", k)
	}

	for k := range seen {
		fmt.Printf("I saw %s\n", k)
	}

	fmt.Println(5 / 2)
}
