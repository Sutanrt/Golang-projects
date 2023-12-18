package main 
import "fmt"

func f(a,b, x int) int {
	var y int
	y=(a*x)+b
	if y<0{
		y=-y
	}
	return y
}

func main () {
	var a,b,k,x int
	fmt.Scan(&a,&b,&k,&x)
	for i:=0;i<k;i++{
		x=f(a,b,x)
	}
	fmt.Print(x)
}