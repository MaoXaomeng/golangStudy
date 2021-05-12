package main

import "fmt"

func main() {
	ages := map[string]int{"alice": 31, "charlie": 34}
	age, ok := ages["bob"]
	fmt.Printf("%d\t%t", age, ok)
	fmt.Printf("%d\n", ages["alice"])
}
