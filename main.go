package main

import "fmt"

func main() {
	valueList := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(sum(valueList...))
}
func sum(vals ...int) int {
	var total int
	for _, val := range vals {
		total += val
	}
	return total
}
