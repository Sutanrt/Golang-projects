package main 
import "fmt"

func proc (x,y float64 ,fx *float64){
	*fx=(2*x)-(y/2)+3
}

func main () {
	var x,y,fx float64
	fmt.Scan(&x,&y)
	proc(x,y,&fx)
	fmt.Print(fx)
}