package main 

import "fmt"

func recursive (x, y float64) float64 {
	if y>=0 || x==0{
	if x==0 {
		return 0
	}else {
		if y==1 {
		return x
	}else if y==0 {
	
		return 1
	}else {
		return x*recursive(x,y-1)
	}
	}
	
	
	}else {
		if y==-1 {
			return 1/x
		}else {
			
			return ((1/x)*(recursive(x,y+1)))
		}
	
	}
	
}
func main () {
	var x,y float64
	fmt.Scan(&x,&y)
	fmt.Print(recursive(x,y))

}