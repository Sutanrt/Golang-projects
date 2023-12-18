package main

import "fmt"

func proc(n int, teks *string) {
	if n == 0 {
		*teks = *teks + " error"
	} else if n == 1 {
		*teks = *teks + " warning"
	} else if n == 2 {
		*teks = *teks + " informasi"
	}
}

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)
	proc(n, &s)
	fmt.Print(s)
}
