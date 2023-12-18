package main

import "fmt"

type ceks struct {
	hari     int
	tangggal int
	waktu    int
}

func main() {
	ar := make([]int, 10, 10)
	var ars []ceks
	var arws [10]ceks
	cek32(&arws)
	ars = make([]ceks, 10)
	fmt.Println(ars)
	fmt.Println(ars[0].hari)
	ar = append(ar, 2)
	ar[0] = 3
	ar[3] = 5
	//ar = append(ar[3:4], 4)
	//fmt.Println(ar)
	fmt.Println(ar, ar[3:], ar[:3])
	fmt.Print(remove(ar))

}
func remove(ar []int) []int {
	ar = append(ar[:3], ar[3+1:]...)
	return ar
}
func cek32(ars *[10]ceks) {

}
