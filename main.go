package main

import (
	"fmt"
)

func main() {
	//	xCapital := make(chan byte)
	x := "hello!"
	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'
			//			xCapital <- x
			fmt.Printf("%c", x)
		}
	}
}
