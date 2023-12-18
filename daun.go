package main 
import "fmt"
func main () {
	var n,lebar,terbesar int
	fmt.Scan(&n)
	terbesar=0
	for i:=0;i<n;i++{
		fmt.Scan(&lebar)
		if terbesar<lebar {
			terbesar=lebar
		}
	}
	fmt.Print(terbesar)
}