package main

import "fmt"

func main() {
	fmt.Println(double(4))
}
func double(x int) (result int) {
	defer func() {
		result += x
		fmt.Printf("double(%d) = %d\n", x, result)
	}()
	result = x + x
	return
}
