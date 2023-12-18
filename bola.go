package main

import (
	"fmt"
)

func main() {
	var n, a, b, c int
	fmt.Scan(&n)
	a = 0
	b = 0
	c = 1
	for i := 0; i < n; i++ {
		fmt.Print(b, " ")
		b = b + c
		c = a
		a = b
	}
}
