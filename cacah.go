package main 
import "fmt"

func main () {
	var x,terbesar int 
	fmt.Scan(&x)
	terbesar=0
	for x>0 {
		if terbesar<x%10 {
			terbesar=x%10
		}
		x=x/10
	}
	fmt.Print(terbesar)

}