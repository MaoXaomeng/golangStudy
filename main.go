package main

import (
	"fmt"
)

func main() {
	f := square
	fmt.Printf("%T\n", f)
	fmt.Println(f(3))
	f = negative
	fmt.Printf("%T\n", f)
	fmt.Printf("%T", product)

}
func square(n int) int     { return n * n }
func negative(n int) int   { return -n }
func product(m, n int) int { return m * n }
