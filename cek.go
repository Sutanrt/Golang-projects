package main

import "fmt"

func main() {
	var c rune
	var x [3]rune
	fmt.Scanf("%c %c", &c, &x[0])

	fmt.Print(c, " ", string(c), " ", c == '.', x, x[0], " ", string(x[0]))

}
