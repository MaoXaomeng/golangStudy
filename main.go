package main

import (
	"fmt"
)

func main() {
	a := [5]int{7, 8, 10}
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	fmt.Printf("hello world")
}
