package main

import "fmt"

func recursive(n int) int {
	if n == 1 {
		fmt.Print(1)
		return 1
	}
	if n%2 != 0 {
		fmt.Print(n, " ")
		return recursive(n - 2)

	} else {
		return recursive(n - 1)
	}

}
func faktor(n, i int) int {
	if n == 1 {
		fmt.Print("1")
		return 1
	} else if i%n == 0 {
		fmt.Print(n, " ")
	}
	//fmt.Println(n, i)
	return (faktor(n-1, i))

}
func pangkat(x, y int) int {
	if y == 0 {
		return 1
	} else if y == 1 {
		return x
	} else {
		return x * pangkat(x, y-1)

	}

}
func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1

	} else {

		return fibonacci(n-1) + fibonacci(n-2)
	}

}
func main() {
	var x int
	fmt.Scan(&x)
	//fmt.Scan(&n)
	//recursive(n)
	//faktor(n, n)
	//fmt.Print(pangkat(x, y))
	fibonacci(x)
}
